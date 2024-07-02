export function Home(userimage,Username,title, content, date,image,like, comment, share) {
    return `<div class="userpost">
                <div class="userimage">
                    <img width="80px" src="${userimage}" alt="">
                    <p>${Username}</p>
                </div>
                <h3>${title}</h3>
                <div class="postcontent">
                    <img  width="100%" src="${image}" alt="" >
                    <p>${content}</p>
                </div>
                <p>${date}</p>
                <div class="line-horizontal"></div>
                <div class="like-comment-share">
                    <div class="num">
                        <div id="likepost"></div>
                        <p> ${like}</p>
                    </div>
                    <div class="num">
                        <div id="comment" class="comment"></div>
                         <p>${comment} </p>
                        
                    </div>
                    <div class="num">
                        <div id="sharepost"></div>
                         <p>${share}</p>
                        
                    </div>
                </div>
            </div>`
}

export function Post(Title, Content, Image, Categorie, Author, userimage, Date) {
    return ` <div class="userpost">
    <div class="userimage"> <img src="${userimage}" alt=""></div>
    <p>date</p>
    <h3>${Title}</h3>
    <div class="postcontent">
        <img src="${Image}" alt="" srcset="">
        <p>${Content}</p>
    </div>
    <p>${Date}/p>
    <div class="like-comment-share">
        <div id="likepost"></div>
        <div id="comment" class="comment"></div>
        <div id="sharepost"></div>
    </div>
</div>`
}

export function RegisterAndLogin(){
    return ` <div class="forms">
        <form id="register-form" method="post" class="hide">
            <div class="form-header">
                <h1>Sign up</h1>
                <p>Please fill in the form below to register.</p>
            </div>

            <div class="form-body">
                <div class="message hide">
                    <span class="error-message">Hello</span>
                    <span class="success-message"></span>
                </div>

                <div class="form-group">
                    <label for="nickname">Nickname</label>
                    <input type="text" id="nickname" name="nickname" required autocomplete="off">
                </div>

                <div class="form-group">
                    <label for="age">Age</label>
                    <input type="number" id="age" name="age" required step="1" min="12" max="120" autocomplete="off">
                </div>

                <div class="form-group">
                    <label for="gender">Gender</label>
                    <select name="gender" id="gender">
                        <option value="none" selected disabled>---- Choose gender ----</option>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                    </select>
                </div>

                <div class="form-group">
                    <label for="firstname">First name</label>
                    <input type="text" id="firstname" name="firstname" required autocomplete="off">
                </div>

                <div class="form-group">
                    <label for="lastname">Last name</label>
                    <input type="text" id="lastname" name="lastname" required autocomplete="off">
                </div>

                <div class="form-group">
                    <label for="email">E-mail</label>
                    <input type="email" id="email" name="email" required autocomplete="off">
                </div>

                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" id="password" name="password" required autocomplete="off">
                </div>

                <div class="form-group" style="margin-top: 12px">
                    <button id="register-button">Register</button>
                </div>
            </div>
            <div class="form-footer">
                <p>Already have an account? <a class="go-login">Login</a></p>
                <a class="go-home">
                    <i class="material-icons">home</i>Home
                </a>
            </div>
        </form>

        <form id="login-form" method="post" class="">
            <div class="form-header">
                <h1>Sign in</h1>
                <p>Please fill in the form below to login.</p>
            </div>
            <div class="form-body">
                <div class="message hide">
                    <span class="error-message hide"></span>
                    <span class="success-message hide"></span>
                </div>

                <div class="form-group">
                    <label for="username">E-mail or nickname</label>
                    <input type="text" id="username" name="nickname" required autocomplete="off">
                </div>

                <div class="form-group">
                    <label for="pass">Password</label>
                    <input type="password" id="pass" name="password" required autocomplete="off">
                </div>

                <div class="form-group" style="margin-top: 12px">
                    <button id="login-button">Login</button>
                </div>
            </div>
            <div class="form-footer">
                <p>Don't have an account? <a class="go-register">Register</a></p>
            </div>
        </form>
    </div>
`
}

export function Messageinterface(){
    return ` <div id="main-message">
        <h2>Chat Messages</h2>
        <div class="all-message">
            <div class="container">
                <img src="/w3images/bandmember.jpg" alt="Avatar" style="width:100%;">
                <p>Hello. How are you today?</p>
                <span class="time-right">11:00</span>
            </div>

            <div class="container darker">
                <img src="/w3images/avatar_g2.jpg" alt="Avatar" class="right" style="width:100%;">
                <p>Hey! I'm fine. Thanks for asking!</p>
                <span class="time-left">11:01</span>
            </div>
        </div>
        <div class="input-message">
            <form action="" id="input-message" method="post">
                <textarea rows="4" cols="70" id="message-send" name="message"></textarea>
                <input type="submit" id="submit" value="Envoyer">
            </form>
        </div>

    </div>`
}


export function PostForm() {
    return `<div class="form">
     <div class="title">
      <h2>New post</h2>
      <div class="barre"></div>
    </div>
      <form id="createPost" action="">
   <div class="input">
    <label for="title">Title</label>
    <input type="text" id="title" name="title" required></div>
    
    <div class="input">
    <label for="content">Description</label>
    <textarea id="description" name="content" rows="4" cols="50" required></textarea>
    </div>
    
    <div id="input-Image" class="input">
        <input type="file" id="fileUpload"  name="image">
        <label for="fileUpload" class="custom-file-upload">
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-file-earmark-plus" viewBox="0 0 16 16">
  <path d="M8 6.5a.5.5 0 0 1 .5.5v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 .5-.5"/>
  <path d="M14 4.5V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h5.5zm-3 0A1.5 1.5 0 0 1 9.5 3V1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4.5z"/>
</svg>
            <span> Choisir un fichier</span>
        </label>
    </div>  

    <div class="input">
        <input type="submit" value="Submit">
     </div>  
</form>
  </div>`
}