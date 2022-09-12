---
title: Gradient API v1.0
language_tabs: []
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="gradient-api">Gradient API v1.0</h1>

> Scroll down for example requests and responses.

Base URLs:

* <a href="http://localhost:3000/">http://localhost:3000/</a>

# Authentication

- HTTP Authentication, scheme: bearer 

<h1 id="gradient-api-default">Default</h1>

## get__

> Code samples

`GET /`

Get a random gradient

> Example responses

> 500 Response

```json
{
  "error": "server error"
}
```

<h3 id="get__-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|success|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<h3 id="get__-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## post__

> Code samples

`POST /`

Get a gradient with specified parameters

> Body parameter

```json
{
  "colors": [
    {
      "r": 255,
      "g": 255,
      "b": 255,
      "a": 255
    },
    {
      "r": 255,
      "g": 255,
      "b": 255,
      "a": 255
    }
  ],
  "x": 64,
  "y": 64
}
```

<h3 id="post__-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gradient](#schemagradient)|false|none|

> Example responses

<h3 id="post__-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|success|None|

<h3 id="post__-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## post__register

> Code samples

`POST /register`

Create a GradientAPI account to assemble your unique collection of gradients

> Body parameter

```json
{
  "username": "o111oo11o",
  "password": "string"
}
```

<h3 id="post__register-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[register_body](#schemaregister_body)|true|none|

> Example responses

> 201 Response

```json
{
  "accessToken": "accessToken",
  "refreshToken": "refreshToken"
}
```

<h3 id="post__register-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|success|[inline_response_201](#schemainline_response_201)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<aside class="success">
This operation does not require authentication
</aside>

## post__login

> Code samples

`POST /login`

Login to GradientAPI

> Body parameter

```json
{
  "username": "o111oo11o",
  "password": "string"
}
```

<h3 id="post__login-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[login_body](#schemalogin_body)|true|none|

> Example responses

> 200 Response

```json
{
  "accessToken": "accessToken",
  "refreshToken": "refreshToken"
}
```

<h3 id="post__login-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|success|[inline_response_201](#schemainline_response_201)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|invalid credentials|[inline_response_422](#schemainline_response_422)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<aside class="success">
This operation does not require authentication
</aside>

## post__refresh

> Code samples

`POST /refresh`

Refresh your access tokens before they expire!

> Body parameter

```json
{
  "refreshToken": "string"
}
```

<h3 id="post__refresh-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[refresh_body](#schemarefresh_body)|true|none|

> Example responses

> 200 Response

```json
{
  "accessToken": "accessToken",
  "refreshToken": "refreshToken"
}
```

<h3 id="post__refresh-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|success|[inline_response_201](#schemainline_response_201)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|invalid credentials|[inline_response_422](#schemainline_response_422)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## post__logout

> Code samples

`POST /logout`

Log out

> Example responses

> 401 Response

```json
{
  "error": "string"
}
```

<h3 id="post__logout-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|you are logged out|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|invalid credentials|[inline_response_422](#schemainline_response_422)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## get__user_{username}

> Code samples

`GET /user/{username}`

get gallery of

<h3 id="get__user_{username}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|path|string|true|none|

> Example responses

> 200 Response

```json
{
  "gradientList": [
    "",
    ""
  ]
}
```

<h3 id="get__user_{username}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|abc|[inline_response_200](#schemainline_response_200)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|invalid credentials|[inline_response_422](#schemainline_response_422)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<aside class="success">
This operation does not require authentication
</aside>

## post__gradient

> Code samples

`POST /gradient`

save gradient to your collection

> Body parameter

```json
{
  "gradient": {
    "colors": [
      {
        "r": 255,
        "g": 255,
        "b": 255,
        "a": 255
      },
      {
        "r": 255,
        "g": 255,
        "b": 255,
        "a": 255
      }
    ],
    "x": 64,
    "y": 64
  },
  "name": "string"
}
```

<h3 id="post__gradient-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[gradient_body](#schemagradient_body)|false|none|

> Example responses

> 401 Response

```json
{
  "error": "string"
}
```

<h3 id="post__gradient-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Gradient saved to your collection|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|invalid credentials|[inline_response_422](#schemainline_response_422)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<h3 id="post__gradient-responseschema">Response Schema</h3>

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## get__gradient_{id}

> Code samples

`GET /gradient/{id}`

<h3 id="get__gradient_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|

> Example responses

> 401 Response

```json
{
  "error": "string"
}
```

<h3 id="get__gradient_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Image with given id|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|invalid credentials|[inline_response_422](#schemainline_response_422)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<h3 id="get__gradient_{id}-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## delete__gradient_{id}

> Code samples

`DELETE /gradient/{id}`

<h3 id="delete__gradient_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|

> Example responses

> 200 Response

```json
{
  "success": "success"
}
```

<h3 id="delete__gradient_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|gradient successfully deleted|[inline_response_200_1](#schemainline_response_200_1)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|invalid credentials|[inline_response_422](#schemainline_response_422)|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|invalid parameters|[inline_response_422](#schemainline_response_422)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|internal server error|[inline_response_500](#schemainline_response_500)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Color">Color</h2>
<!-- backwards compatibility -->
<a id="schemacolor"></a>
<a id="schema_Color"></a>
<a id="tocScolor"></a>
<a id="tocscolor"></a>

```json
{
  "r": 255,
  "g": 255,
  "b": 255,
  "a": 255
}

```

color data type

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|r|integer|true|none|none|
|g|integer|true|none|none|
|b|integer|true|none|none|
|a|integer|false|none|none|

<h2 id="tocS_Gradient">Gradient</h2>
<!-- backwards compatibility -->
<a id="schemagradient"></a>
<a id="schema_Gradient"></a>
<a id="tocSgradient"></a>
<a id="tocsgradient"></a>

```json
{
  "colors": [
    {
      "r": 255,
      "g": 255,
      "b": 255,
      "a": 255
    },
    {
      "r": 255,
      "g": 255,
      "b": 255,
      "a": 255
    }
  ],
  "x": 64,
  "y": 64
}

```

color data type

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|colors|[[Color](#schemacolor)]|false|none|[color data type]|
|x|number|false|none|none|
|y|number|false|none|none|

<h2 id="tocS_inline_response_500">inline_response_500</h2>
<!-- backwards compatibility -->
<a id="schemainline_response_500"></a>
<a id="schema_inline_response_500"></a>
<a id="tocSinline_response_500"></a>
<a id="tocsinline_response_500"></a>

```json
{
  "error": "server error"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|error|string|false|none|none|

<h2 id="tocS_register_body">register_body</h2>
<!-- backwards compatibility -->
<a id="schemaregister_body"></a>
<a id="schema_register_body"></a>
<a id="tocSregister_body"></a>
<a id="tocsregister_body"></a>

```json
{
  "username": "o111oo11o",
  "password": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|false|none|none|
|password|string|false|none|none|

<h2 id="tocS_inline_response_201">inline_response_201</h2>
<!-- backwards compatibility -->
<a id="schemainline_response_201"></a>
<a id="schema_inline_response_201"></a>
<a id="tocSinline_response_201"></a>
<a id="tocsinline_response_201"></a>

```json
{
  "accessToken": "accessToken",
  "refreshToken": "refreshToken"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|accessToken|string|false|none|none|
|refreshToken|string|false|none|none|

<h2 id="tocS_inline_response_422">inline_response_422</h2>
<!-- backwards compatibility -->
<a id="schemainline_response_422"></a>
<a id="schema_inline_response_422"></a>
<a id="tocSinline_response_422"></a>
<a id="tocsinline_response_422"></a>

```json
{
  "error": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|error|string|false|none|none|

<h2 id="tocS_login_body">login_body</h2>
<!-- backwards compatibility -->
<a id="schemalogin_body"></a>
<a id="schema_login_body"></a>
<a id="tocSlogin_body"></a>
<a id="tocslogin_body"></a>

```json
{
  "username": "o111oo11o",
  "password": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|false|none|none|
|password|string|false|none|none|

<h2 id="tocS_refresh_body">refresh_body</h2>
<!-- backwards compatibility -->
<a id="schemarefresh_body"></a>
<a id="schema_refresh_body"></a>
<a id="tocSrefresh_body"></a>
<a id="tocsrefresh_body"></a>

```json
{
  "refreshToken": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|refreshToken|string|false|none|none|

<h2 id="tocS_inline_response_200">inline_response_200</h2>
<!-- backwards compatibility -->
<a id="schemainline_response_200"></a>
<a id="schema_inline_response_200"></a>
<a id="tocSinline_response_200"></a>
<a id="tocsinline_response_200"></a>

```json
{
  "gradientList": [
    "",
    ""
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|gradientList|[any]|false|none|none|

<h2 id="tocS_gradient_body">gradient_body</h2>
<!-- backwards compatibility -->
<a id="schemagradient_body"></a>
<a id="schema_gradient_body"></a>
<a id="tocSgradient_body"></a>
<a id="tocsgradient_body"></a>

```json
{
  "gradient": {
    "colors": [
      {
        "r": 255,
        "g": 255,
        "b": 255,
        "a": 255
      },
      {
        "r": 255,
        "g": 255,
        "b": 255,
        "a": 255
      }
    ],
    "x": 64,
    "y": 64
  },
  "name": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|gradient|[Gradient](#schemagradient)|false|none|color data type|
|name|string|true|none|none|

<h2 id="tocS_inline_response_200_1">inline_response_200_1</h2>
<!-- backwards compatibility -->
<a id="schemainline_response_200_1"></a>
<a id="schema_inline_response_200_1"></a>
<a id="tocSinline_response_200_1"></a>
<a id="tocsinline_response_200_1"></a>

```json
{
  "success": "success"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|success|string|false|none|none|

