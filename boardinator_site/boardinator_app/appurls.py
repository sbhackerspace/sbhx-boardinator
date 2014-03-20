from django.conf.urls import url, patterns
from django.contrib.auth.views import login, logout

urlpatterns = patterns('boardinator_app.views',
    url(r'^$',    			'base', 			name='base'),
    # url(r'^task/', 			'task', 			name='task'),
    # url(r'^calendar/', 	'calendar', 	name='calendar'),
)




