
    <div class="col-md-12">
        <div class="avatar pull-left">
            <img src="{{asset($comment->owner->avatar)}}" class="img-bordered img-responsive"
                 style="width: 60px;height: 60px;">
        </div>
        <div class="c_k">
            @if($comment->co_pid==0)
                <span style="color:#31b0d5">{{$comment->owner->name}}</span>:
                <span>{{$comment->co_content or ''}}</span><br/>
            @else
                <span style="color:#31b0d5">{{$comment->owner->name}}&nbsp;回复&nbsp;{{$comment->reply_name}}</span>:
                <span>{{$comment->co_content or ''}}</span><br/>
            @endif
            <span style="width: 60px;">&emsp;&emsp;&emsp;&emsp;</span>
            <span style="color: red">{{$comment->created_at or ''}}</span>&nbsp;
            <a class="c_reply">回复</a>

            @include('home.comments.form',['parentId'=>$comment->co_id])

            @if(isset($a_comment[$comment->co_id]))
                @include('home.comments.list',['collections'=>$a_comment[$comment->co_id]])
            @endif

            <div class="line"></div>
        </div>
    </div>


