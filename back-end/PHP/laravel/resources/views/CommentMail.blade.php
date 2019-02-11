@component('mail::message')
    {{$username}}在文章<<{{$title}}>>中评论了。<br/>
    评论内容：{{$content}}

@component('mail::button', ['url' => $url])
    {{$title}}
@endcomponent

Thanks,<br>
[{{config('app.name')}}]({{config('blog.blog_home')}})
@endcomponent
