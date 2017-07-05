<!DOCTYPE html>
<html>
	<head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
        <link rel="shortcut icon" href="/static/img/favicon.png" />
        <meta name="author" content="Unknown" />
        <meta name="description" content="Beego sample app - Web IM" />
        <meta name="keywords" content="Go, golang, beego, sample, Web IM">
        <link href="/static/css/bootstrap.min.css" rel="stylesheet" />
        <script src="/static/js/jquery-1.10.1.min.js"></script>
    </head>
	<body>
            <div class="container" style="width: 500px;">
                <h3>hhhhh</h3>
                <br>
                <form action="/join" method="post" class="form-horizontal">
                    <div class="form-group">
                        <label class="col-md-3 control-label">1111: </label>
                        <div class="col-md-5">
                              <input type="text" class="form-control" name="uname" required>
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-md-3 control-label">2222222222: </label>
                        <div class="col-md-5">
                            <select class="form-control" name="tech">
                                <option value="longpolling">3333333</option>
                                <option value="websocket">44444444444444</option>
                            </select>
                        </div>
                    </div>

                    <div class="form-group">
                        <div class="col-sm-offset-3 col-sm-10">
                            <button type="submit" class="btn btn-info">ssss</button>
                        </div>
                    </div>
                </form>
            </div>
	        hello {{.name}} , your age is {{.age}} !
	</body>
</html>