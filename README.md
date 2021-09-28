# short-url

requirements:
- shorten long url into a 8-character string
- avoid collision

outcome:
- deploy this to the main site
- take in random urls and return shortened url
- able to copy on the result directly

backend:
- shortening
- store the mapping <longUrl, shortUrl>
- able to re-direct

frontend:
- consume this backend service
