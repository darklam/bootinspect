### Running

`go run main.go --file app.css`

### Example

Taking a look at this code,

```css
.foo {
    text-transform: uppercase;
    height: 100px;
    width: 100%;
}
```

and assuming we have bootstrap.css in our project's dependencies, we see that text-transform and width can be replaced with bootstrap's classes `text-uppercase` and `w-100`

so instead of writting extra CSS, you could just do that

```html
<div class="foo text-uppercase w-100">
</div>
```

Notice that the foo is still used, that's because `height: 100px` is **not** available on bootstrap's css utilities.