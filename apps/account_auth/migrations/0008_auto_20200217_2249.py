# Generated by Django 2.2.5 on 2020-02-17 14:49

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('account_auth', '0007_auto_20200217_1544'),
    ]

    operations = [
        migrations.AlterField(
            model_name='userprofile',
            name='user_img',
            field=models.ImageField(default=None, upload_to='userprofile/headpicture/'),
        ),
    ]