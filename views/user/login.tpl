<form class="form-signin" method="POST" >
      <img class="mb-4" src="https://getbootstrap.com/assets/brand/bootstrap-solid.svg" alt="" width="72" height="72">
      <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1>
      <input type="text" id="username" class="form-control" placeholder="username" required autofocus>
      <input type="password" id="password" class="form-control" placeholder="Password" required>
      <div class="checkbox mb-3">
        <label>
          <input type="checkbox" value="remember-me"> Remember me
        </label>
      </div>
      <button class="btn_submit btn btn-lg btn-primary btn-block" type="button">Sign in</button>
      <p class="mt-5 mb-3 text-muted">&copy; 2017-2018</p>
</form>


<script>
    $(function(){

     $(".btn_submit").on('click' ,function(){
      var username = $("#username").val();
      var password = $("#password").val();
      $.ajax({
        type : "POST",
        url  : "{{.formSubmitUrl}}",
        async : false,
        data : "name=" + username + "&pwd=" + password,
        dataType:"json",
        success: function (data) {
          console.log(data)
            if(data.code == "0"){
                window.location.href = "{{.homeUrl}}"
            }else{
              alert(data.msg)
              return
            }
        }
      });
    });
 })

</script>