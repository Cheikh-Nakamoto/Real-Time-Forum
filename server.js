const http = require('http');
const fs = require('fs');
const path = require('path');

let PORT = process.env.PORT || 4000

const getMimeType = (ext) => {
    const mimeTypes = {
        '.html': 'text/html',
        '.css': 'text/css',
        '.js': 'application/javascript',
        '.png': 'image/png',
        '.jpg': 'image/jpeg',
        '.gif': 'image/gif',
        '.svg': 'image/svg+xml',
        '.json': 'application/json',
        '.woff': 'font/woff',
        '.woff2': 'font/woff2',
        '.ttf': 'font/ttf',
        '.eot': 'application/vnd.ms-fontobject',
        '.otf': 'font/otf',
        '.wasm': 'application/wasm'
    };
    return mimeTypes[ext] || 'application/octet-stream';
};

let server = http.createServer((req, res) => {
    let filePath = '';
    // Définir le chemin du fichier basé sur la route
    if (req.url === '/') {
        filePath = path.join(__dirname, 'assets', 'html', 'index.html');
    } else if (req.url.endsWith('.css')) {
        filePath = path.join(__dirname, req.url);
    } else if (req.url.endsWith('.js')) {
        filePath = path.join(__dirname,  req.url);
    } else if (req.url.endsWith('.html')) {
        filePath = path.join(__dirname,req.url);
    } else if (req.url.endsWith('.json')) {
        let subpath = req.url.split('/');
        filePath = path.join(__dirname,'assets' , subpath[2], subpath[3])
    }else {
        filePath = path.join(__dirname,req.url);;
    }

    const ext = path.extname(filePath);
    const contentType = getMimeType(ext);
    //const filepath = path.join(__dirname,'index.html')
    fs.readFile(filePath, (err, content) => {
        if (err) {
            if (err.code === 'ENOENT') {
                // Si le fichier n'est pas trouvé, renvoyer une réponse 404
                res.writeHead(404, { 'Content-Type': 'text/html' });
                res.end('<h1>404 - Page Not Found</h1>', 'utf8');
            } else {
                // Pour d'autres erreurs, renvoyer une réponse 500
                res.writeHead(500);
                res.end(`Server Error: ${err.code}`);
            }
        } else {
            // Renvoyer le contenu du fichier avec le type MIME approprié
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content, 'utf8');
        }
    });
    
})

server.listen(PORT, () => {
    console.log(`Server is running on port: http://localhost:${PORT}`)
})