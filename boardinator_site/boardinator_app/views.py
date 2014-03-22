from django import forms
from django.contrib import messages
from django.core import serializers
from django.core.context_processors import csrf
from django.http import HttpResponse, HttpResponseRedirect
from django.shortcuts import render_to_response, get_object_or_404, render, \
    redirect
from django.template import loader, RequestContext
from django.views.decorators.csrf import csrf_exempt

def index(request):
    return render(request, "index.html", locals())

def task(request):
		return render(request, "task.html", locals())

def calendar(request):
		return render(request, "calendar.html", locals())

def ang(request):
	return render(request, "app/index.html")
		

