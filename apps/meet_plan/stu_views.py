from django.core.exceptions import PermissionDenied
from django.shortcuts import get_object_or_404
from django.urls import reverse

from django.views.generic import ListView
from django.db.models import Count, Sum, When, Case, BooleanField
from django.db.models import Q
from django.utils import timezone
from django.views.generic.edit import CreateView

from utils.mixin.permission import LoginRequiredMixin, UserProfileRequiredMixin

from ..account_auth.models import User
from .models import MeetPlan, MeetPlanOrder
from .forms import MeetPlanOrderCreateForm
from .utils import get_term_date


class TeacherListView(LoginRequiredMixin, UserProfileRequiredMixin, ListView):
    template_name = 'meet_plan/student/teacher_all.html'
    context_object_name = 'teacher_list'

    def get_queryset(self):
        date_range = get_term_date()

        queryset = User.objects.filter(is_teacher=True) \
            .annotate(meetplan_num=Count('meetplan',
                                         filter=Q(meetplan__start_time__gt=date_range[0],
                                                  meetplan__end_time__lt=date_range[1]),
                                         distinct=True
                                         ),
                      meetplan_available_num=Count('meetplan',
                                                   filter=Q(meetplan__start_time__gt=timezone.now(),
                                                            meetplan__end_time__lt=date_range[1],
                                                            meetplan__available_choice__gt=0),
                                                   distinct=True
                                                   ),
                      meetplanorder_available_num=Sum('meetplan__available_choice',
                                                      filter=Q(meetplan__start_time__gt=timezone.now(),
                                                               meetplan__end_time__lt=date_range[1])
                                                      )
                      ) \
            .order_by('user_name')
        return queryset


class TeacherPlanListView(LoginRequiredMixin, UserProfileRequiredMixin, ListView):
    template_name = 'meet_plan/student/teacher_plan_all.html'
    context_object_name = 'plan_list'
    paginate_by = 50
    teacher = ''

    def get_queryset(self):
        date_range = get_term_date()
        self.teacher = get_object_or_404(User, pk=self.kwargs['pk'])
        if not self.teacher.is_teacher:
            raise PermissionDenied('您将要查看的用户身份为学生，这是不被允许的！')
        return MeetPlan.objects.filter(teacher=self.teacher,
                                       start_time__gt=date_range[0],
                                       end_time__lt=date_range[1]) \
            .annotate(available=Case(When(start_time__lt=timezone.now(), then=False),
                                     default=True,
                                     output_field=BooleanField())) \
            .order_by('start_time')

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['teacher_id'] = self.teacher.pk
        context['teacher_name'] = self.teacher.user_name
        return context


class MeetPlanOrderCreateView(LoginRequiredMixin, UserProfileRequiredMixin, CreateView):
    model = MeetPlanOrder
    template_name = 'meet_plan/student/order_create.html'
    form_class = MeetPlanOrderCreateForm
    meet_plan = ''

    def form_valid(self, form):
        form.instance.meet_plan = self.meet_plan
        form.instance.student = self.request.user
        response = super().form_valid(form)

        from .tasks import send_meetplan_order_create_email
        domain = self.request.get_host()
        send_meetplan_order_create_email.delay(self.object.id, domain)
        return response

    def post(self, request, *args, **kwargs):
        date_range = get_term_date()
        self.meet_plan = get_object_or_404(MeetPlan, pk=self.kwargs['pk'])
        if MeetPlanOrder.objects.filter(student=self.request.user,
                                        meet_plan__start_time__gt=date_range[0],
                                        meet_plan__end_time__lt=date_range[1]
                                        ).count() >= 2:
            raise PermissionDenied('您本学期已达2次选课上限！')
        if not self.meet_plan.available_choice or self.meet_plan.start_time < timezone.now():
            raise PermissionDenied('该安排已满或该安排已过期！')

        if MeetPlanOrder.objects.filter(student=self.request.user,
                                        meet_plan=self.meet_plan).count() > 0:
            return PermissionDenied('您不能多次选同一个综合指导课安排！')

        return super().post(request, *args, **kwargs)

    def get_success_url(self):
        return reverse('meet_plan:index')