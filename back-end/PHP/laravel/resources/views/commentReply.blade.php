@component('mail::message')
# {{$reply_name}} 你好

{{$comment_name}}在文章<<{{$title}}>>中回复了你:

{{$reply_content}}

Thanks,<br>
[{{ config('app.name') }}]({{config('blog.blog_home')}})
@endcomponent
