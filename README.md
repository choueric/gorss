# Intro
GoRSS is a RSS feed generator for gzh(public account) of weixin.
It runs on Google App Engine.

# Usage
Modify input.go to add new gzh feed. But before that, you should fetch some
information about the gzh by yourself manually, such as openID, eqs, etc. 

This job will be added to GoRSS later, someday, maybe, who knowns.

After adding new feed and deployint to GAE, you can see the feed in the URL:

```
http://appid.appspot.com/gzhid_rss
```

In this URL, appid is your application id. And gzhid is the id of the gzh which
should be as same as the one you add in input.go file.
