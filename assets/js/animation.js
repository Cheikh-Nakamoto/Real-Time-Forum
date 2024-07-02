export function Animation() {
    function loadLottieAnimation(element, path) {
        return lottie.loadAnimation({
            container: element,
            renderer: 'svg', // Choisissez le rendu SVG ou canvas
            loop: true, // Définissez si l'animation doit boucler
            autoplay: false, // Ne démarrez pas automatiquement l'animation
            path: path // Chemin vers votre fichier JSON LottieFlow
        });
    }

    // Fonction pour ajouter des écouteurs d'événements pour démarrer et arrêter l'animation
    function addHoverAnimation(element, animation) {
        element.addEventListener('mouseenter', function () {
            animation.play();
        });
        element.addEventListener('mouseleave', function () {
            animation.stop();
        });
    }

    // Charger et ajouter des animations pour chaque élément
    document.querySelectorAll('#like').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/like.json');

        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#communaute').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/groupe.json');
        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#message').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/message.json');
        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#share').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/share.json');
        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#post').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/post.json');
        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#comment').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/commentaire.json');
        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#sharepost').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/share.json');
        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#likepost').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/like.json');
        addHoverAnimation(element, animation);
    });

    document.querySelectorAll('#notification').forEach(function (element) {
        element.innerHTML = '';
        const animation = loadLottieAnimation(element, '/static/js/notification.json');
        addHoverAnimation(element, animation);
    });
}



