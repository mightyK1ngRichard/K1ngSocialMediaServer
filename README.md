# Backend Server for K1ngSocialMedia

[ iOS-application](https://github.com/mightyK1ngRichard/K1ngSocialMedia)

```http
GET /users                   # Пользователи
GET /users?id=               # Пользователь
GET /posts?user_id=          # Посты пользователя
GET /comments?user_id=       # Комментарии под постами пользователя
POST /user/{user_id}/upload  # Загрузить фото пользователя
```

### Пример post запроса:

```curl
➜  ~ curl -X POST http://localhost:8010/user/1/upload \
  -F "file=@/Users/dmitriy/Downloads/SpiderImages/iu3.jpg;type=image/jpeg"

# Ответ:
"./static/img/1688243307056518000_13.jpg"% 
```