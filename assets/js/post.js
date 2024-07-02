import { Post, PostForm } from '/assets/js/compenant.js'

let posts = document.getElementById("posts")

posts.addEventListener('click', (e) => {
    let token = localStorage.getItem("token")
    if (location.pathname != "/post" && token) {
        window.history.pushState(null, null, '/post')
        console.log(location.pathname)
        document.getElementById("change").innerHTML = PostForm()
        let form = document.getElementById("createPost")
        form.addEventListener('submit', (e) => {
            e.preventDefault()
            const formdata = new FormData(form)
            const data = Object.fromEntries(formdata.entries())
            console.log("Ces donne on ete envoyer", data, formdata)
            fetch('http://localhost:1112/api/posts', {
                method: 'POST',
                body: formdata,
            }).then((res) => res.json())
                .then((data) => {

                }).catch((err) => {
                    console.log(err)
                })
        })
    }
})