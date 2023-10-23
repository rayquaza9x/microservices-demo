curl --request POST \
  --url http://localhost:3000/auth/register \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "elon@musk.com",
 "password": "1234567"
}'

curl --request POST \
  --url http://localhost:3000/auth/login \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "elon@musk.com",
 "password": "1234567"
}'

curl --request POST \
  --url http://localhost:3000/product/ \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY1MjQ5MzUsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.HBXgv1J7RiOBsHndMHFNcETtGcW9517MRsXK33xRmTE' \
  --header 'Content-Type: application/json' \
  --data '{
 "name": "Product A",
 "stock": 5,
 "price": 15
}'

curl --request POST \
  --url http://localhost:3000/product \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY1MjQ5MzUsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.HBXgv1J7RiOBsHndMHFNcETtGcW9517MRsXK33xRmTE' \
  --header 'Content-Type: application/json' \
  --data '{
 "name": "Product B",
 "stock": 50,
 "price": 150
}'

curl --request GET \
  --url http://localhost:3000/product/1 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY1MjQ5MzUsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.HBXgv1J7RiOBsHndMHFNcETtGcW9517MRsXK33xRmTE'