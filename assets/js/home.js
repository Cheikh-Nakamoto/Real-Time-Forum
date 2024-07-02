import { Home } from '/assets/js/compenant.js'
import { Animation } from '/assets/js/animation.js'
let Communaute = document.getElementById('communaut')
let chat = document.getElementById('message')
let change = document.getElementById('change')

const router = async () => {
    const path = [
        { pathname: '/', component: () => console.log("ome ma nigeu !!") },
        { pathname: '/post', component: () => console.log("post ma nigeu !!") },
        { pathname: '/groupe', component: () => console.log("groupe ma nigeu !!") },
    ]

    const PotentielMatched = path.map((route) => {
        return {
            route: route,
            ismatched: route.pathname == location.pathname
        }
    })

    let match = PotentielMatched.find((route) => route.ismatched)
    if (!match) return {
        path: path[0],
        ismatched: true
    }
}

const navigateTo = (path) => {
    window.history.pushState(null, null, path)
    router()
    
}

document.addEventListener('DOMContentLoaded', (e) => {
    document.body.addEventListener('click', (e) => {
        navigateTo(e.target.href)
    })
    router()

    FetchAllPosts().then((data) => {
        DispatchePost(data)
        Animation()
    })
})


Communaute.addEventListener('click', () => {
    if (location.pathname != "/") {
        navigateTo('/')
        change.innerHTML = ""
        FetchAllPosts().then((data) => {
            DispatchePost(data)
            Animation()
        })
    }
})

chat.addEventListener('click', (e) =>{
    let token = localStorage.getItem('token')
    let ID = localStorage.getItem("object")
    if (token){
        const res = fetch('http://localhost:1112/api/messagerie',{
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({"ID":ID})}
        )
        res.then(res => res.json).then(data => {

        })
        res.catch(err => console.log(err))
    }
})



export const FetchAllPosts = async () => {
    const res = await fetch('http://localhost:1112/api/posts')
    const data = res.json()
    return data
}

export const DispatchePost = (data) => {
    data.forEach(element => {
        const post = Home("https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png", `User ${element.ID}`, element.Title, element.Content, "1H", element.Image, "0", "0", "0")
        change.insertAdjacentHTML('beforeend', post)
    });
}

// export function Home(userimage,Username,title, content, date,image,like, comment, share) {
