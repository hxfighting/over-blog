@component('mail::message')
# {{$name}} 你好

{{$reply_content}}

Thanks,<br>
[{{ config('app.name') }}]({{config('blog.blog_home')}})
@endcomponent
