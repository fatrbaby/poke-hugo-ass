# Poke hugo ass

rebuild markdown files from a [hugo](https://github.com/gohugoio/hugo) site

## build
```shell script
cd bin
go build
```

## config

```yaml
document:
  draft: false #is draft
dir:
  html: /Users/fatrbaby/Charge/blog/public/posts # the hugo generated html directory
  target: /Users/fatrbaby/Charge/blog/content/posts # rebuilt markdown files directory
pattern:
  title: "h2>a" # title find pattern, use goquery library, pattern like jQuery
  date: "span.date"
  content: "div.post_content"
  tags: "meta[name=keywords]"
```

## dependencies
- [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
- [JohannesKaufmann/html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown)

