document.body.onselectstart = () => false
document.body.addEventListener("contextmenu", (e) => {e.preventDefault()})

document.querySelector(".brand").addEventListener('click', () => {
    window.location.href = '/'
})