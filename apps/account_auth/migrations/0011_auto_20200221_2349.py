# Generated by Django 2.2.5 on 2020-02-21 15:49

from django.db import migrations
import phonenumber_field.modelfields


class Migration(migrations.Migration):

    dependencies = [
        ('account_auth', '0010_auto_20200220_0024'),
    ]

    operations = [
        migrations.AlterField(
            model_name='userprofile',
            name='telephone',
            field=phonenumber_field.modelfields.PhoneNumberField(max_length=128, region='CN'),
        ),
    ]