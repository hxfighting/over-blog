@component('mail::message')
    # {{$reply_name}} 你好

    {{$comment_name}}在文章<<{{$title}}>>中回复了你:

    {{$reply_content}}

@component('mail::button', ['url' => $url])
   {{$title}}
@endcomponent

Thanks,<br>
[{{config('app.name')}}]({{config('blog.blog_home')}})
@endcomponent
