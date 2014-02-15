from django.conf.urls import url, patterns, include
from django.contrib import admin
from django.views.generic import TemplateView
from django.conf import settings

admin.autodiscover()

urlpatterns = patterns('',
    ## App URL include
    url(r'^', include('boardinator_app.appurls')),

    ## Static Links
    #url(r'^', TemplateView.as_view(template_name='home.html'), name='home'),

)

if settings.DEBUG:
    urlpatterns = patterns('',
    url(r'^media/(?P<path>.*)$', 'django.views.static.serve',
        {'document_root': settings.MEDIA_ROOT, 'show_indexes': True}),
    url(r'', include('django.contrib.staticfiles.urls')),
) + urlpatterns
