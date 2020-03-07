from django.urls import reverse
from django.views.generic.list import ListView
from django.views.generic.edit import CreateView, UpdateView

from utils.mixin.permission import AdminRequiredMixin
from apps.account_auth.models import User
from utils.mixin.view import FileUploadViewMixin

from .forms import UserCreateForm, UserUpdateForm
from .tasks import send_account_active_email, create_many_user


class UserView(AdminRequiredMixin, ListView):
    model = User
    template_name = 'cmsadmin/user/user_all.html'
    paginate_by = 50
    context_object_name = 'user_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=False).order_by('-identity_id')


class UserCreateView(AdminRequiredMixin, CreateView):
    model = User
    template_name = 'cmsadmin/user/user_create.html'
    form_class = UserCreateForm
    # success_url = '/cmsadmin/user_all/'

    def get_success_url(self):
        return reverse('cmsadmin:user_all')

    def form_valid(self, form):
        response = super().form_valid(form)
        domain = self.request.get_host()
        # 发邮件
        send_account_active_email.delay(self.object.identity_id, domain)
        return response


class CreateManyUserView(AdminRequiredMixin, FileUploadViewMixin):
    template_name = 'cmsadmin/user/user_create_many.html'

    def get_success_url(self):
        return reverse('cmsadmin:user_all')

    def form_valid(self, form):
        response = super().form_valid(form)
        domain = self.request.get_host()
        # 创建任务
        create_many_user.delay(self.object.id, domain)
        return response


class UpdateUserView(AdminRequiredMixin, UpdateView):
    model = User
    form_class = UserUpdateForm
    template_name = 'cmsadmin/user/user_update.html'
    # success_url = '/cmsadmin/user_all/'

    def get_success_url(self):
        return reverse('cmsadmin:user_all')


class DeletedUserListView(AdminRequiredMixin, ListView):
    model = User
    template_name = 'cmsadmin/user/user_deletelist.html'
    paginate_by = 50
    context_object_name = 'user_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=True).order_by('-update_time')
