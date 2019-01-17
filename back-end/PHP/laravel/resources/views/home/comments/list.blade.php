@foreach($collections as $comment)
@include('home.comments.comment',['comment'=>$comment])
@endforeach