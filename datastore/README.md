To create an articles topic:
gcloud pubsub topics create articles

Publish an article:

gcloud pubsub topics publish articles --message '{
  "author": "Isaac Asimov",
  "title": "The Last Question",
  "url": "http://www.multivax.com/last_question.html"
}'


