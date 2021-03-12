# Generated by Django 3.1 on 2020-11-13 16:35

from django.conf import settings
import django.core.validators
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='User',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('password', models.CharField(max_length=128, verbose_name='password')),
                ('last_login', models.DateTimeField(blank=True, null=True, verbose_name='last login')),
                ('is_superuser', models.BooleanField(default=False, help_text='Designates that this user has all permissions without explicitly assigning them.', verbose_name='superuser status')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('identity_id', models.CharField(db_index=True, max_length=16, unique=True, validators=[django.core.validators.RegexValidator(message='账号格式错误。', regex='^\\d{8,12}$')], verbose_name='职工号\\学号')),
                ('user_name', models.CharField(max_length=100, verbose_name='姓名')),
                ('email', models.EmailField(blank=True, max_length=254, null=True, verbose_name='电子邮件')),
                ('is_active', models.BooleanField(default=True, verbose_name='是否激活')),
                ('is_teacher', models.BooleanField(default=False, verbose_name='是否为教职工')),
                ('is_admin', models.BooleanField(default=False, verbose_name='是否为管理员, 管理员可登陆cmsadmin管理页面')),
            ],
            options={
                'verbose_name': '用户',
                'verbose_name_plural': '用户',
            },
        ),
        migrations.CreateModel(
            name='BaseProfile',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('gender', models.BooleanField(choices=[(False, '男'), (True, '女')], default=False, verbose_name='性别')),
                ('birth', models.DateField(blank=True, null=True, verbose_name='生日')),
            ],
            options={
                'verbose_name': '基本信息',
                'verbose_name_plural': '基本信息',
            },
        ),
        migrations.CreateModel(
            name='Department',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('department', models.CharField(max_length=100, unique=True, verbose_name='系所')),
            ],
            options={
                'verbose_name': '系所\\办公室',
                'verbose_name_plural': '系所\\办公室',
            },
        ),
        migrations.CreateModel(
            name='Grade',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('grade', models.CharField(max_length=4, verbose_name='年级')),
            ],
            options={
                'verbose_name': '年级',
                'verbose_name_plural': '年级',
            },
        ),
        migrations.CreateModel(
            name='Major',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('major', models.CharField(max_length=100, verbose_name='专业')),
                ('department', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='account_auth.department', verbose_name='系所')),
            ],
            options={
                'verbose_name': '专业',
                'verbose_name_plural': '专业',
            },
        ),
        migrations.CreateModel(
            name='TeacherProfile',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('phone_number', models.CharField(max_length=17, validators=[django.core.validators.RegexValidator(message='号码格式错误。座机请加区号', regex='^\\+?1?\\d{9,15}$')])),
                ('office', models.CharField(max_length=100, verbose_name='办公室')),
                ('introduce', models.TextField(blank=True, null=True, verbose_name='个人简介')),
                ('department', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='account_auth.department', verbose_name='系所')),
                ('user', models.OneToOneField(on_delete=django.db.models.deletion.DO_NOTHING, to=settings.AUTH_USER_MODEL, verbose_name='教师')),
            ],
            options={
                'verbose_name': '教职工详细信息',
                'verbose_name_plural': '教职工详细信息',
            },
        ),
        migrations.CreateModel(
            name='StudentProfile',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('is_graduate', models.BooleanField(choices=[(True, '研究生'), (False, '本科生')], verbose_name='研究生\\本科生')),
                ('phone_number', models.CharField(max_length=17, validators=[django.core.validators.RegexValidator(message='号码格式错误。', regex='^\\+?1?\\d{9,15}$')])),
                ('dorm', models.CharField(max_length=32, verbose_name='宿舍')),
                ('department', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='account_auth.department', verbose_name='系所')),
                ('grade', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='account_auth.grade', verbose_name='年级')),
                ('major', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='account_auth.major', verbose_name='专业')),
                ('user', models.OneToOneField(on_delete=django.db.models.deletion.DO_NOTHING, to=settings.AUTH_USER_MODEL, verbose_name='学生')),
            ],
            options={
                'verbose_name': '学生详细信息',
                'verbose_name_plural': '学生详细信息',
            },
        ),
    ]