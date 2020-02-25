# Generated by Django 2.2.5 on 2020-02-22 08:56

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('meet_plan', '0005_auto_20200217_1656'),
    ]

    operations = [
        migrations.CreateModel(
            name='SemesterDateRange',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('开始时间', models.DateField(verbose_name='学期开始时间')),
                ('结束时间', models.DateField(verbose_name='学期结束时间')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.AlterField(
            model_name='meetplan',
            name='teacher',
            field=models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to=settings.AUTH_USER_MODEL),
        ),
        migrations.AlterField(
            model_name='meetplanorder',
            name='meet_plan',
            field=models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to='meet_plan.MeetPlan'),
        ),
    ]