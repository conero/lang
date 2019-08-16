client_id                   客户端id


authorization_code          授权码
access_token                可请求 token


token                       具有期限，可自动刷新
refresh_token               刷新令牌


头部传递参数


URL 地址：
    /authorize      授权请求    client_id，redirect_uri 
                    redirect_uri?code=授权码
                    redirect_uri: http://example.com/cb#access_token=2YotnFZFEjr1zCsicMWpAA&state=xyz&token_type=example&expires_in=3600
    /token  
        grant_type =
            'refresh_token 刷新令牌'
        refresh_token=
        client_id     = 客服端id
        client_secret = 客服端id

        相应=>
            {
                "access_token":"2YotnFZFEjr1zCsicMWpAA",
                "token_type":"example",
                "expires_in":3600,
                "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
                "example_parameter":"example_value"
            }