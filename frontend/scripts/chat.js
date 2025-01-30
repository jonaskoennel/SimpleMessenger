
function appendReceivedMessage(username,  message, time) {
    var messageTemplate = `<li class="d-flex justify-content-between mb-4">
            <img src="https://mdbcdn.b-cdn.net/img/Photos/Avatars/avatar-6.webp" alt="avatar" class="rounded-circle d-flex align-self-start me-3 shadow-1-strong" width="60">
            <div class="card">
              <div class="card-header d-flex justify-content-between p-3">
                <p class="fw-bold mb-0">${username}</p>
                <p class="text-muted small mb-0"><i class="far fa-clock"></i>${time}</p>
              </div>
              <div class="card-body">
                <p class="mb-0">
                  ${message}
                </p>
              </div>
            </div>
          </li>`

        var element = document.getElementsByClassName("card-body")[0]
        element.append(messageTemplate)
}

function appendSendMessage(username,  message, time) {
    var messageTemplate = `<li class="p-2 border-bottom">
                <a href="#!" class="d-flex justify-content-between">
                  <div class="d-flex flex-row">
                    <img src="https://mdbcdn.b-cdn.net/img/Photos/Avatars/avatar-1.webp" alt="avatar" class="rounded-circle d-flex align-self-center me-3 shadow-1-strong" width="60">
                    <div class="pt-1">
                      <p class="fw-bold mb-0">${username}</p>
                      <p class="small text-muted">${message}</p>
                    </div>
                  </div>
                  <div class="pt-1">
                    <p class="small text-muted mb-1">${time}</p>
                  </div>
                </a>
              </li>`

    var element = document.getElementsByClassName("card-body")[0]
    element.append(messageTemplate)
}

function loadChat() {
    var username = "Jonas Koennel"
    var message = "Das ist eine neue Nachricht"
    var time = "Just now"
    appendReceivedMessage(username, message, time)
}

window.onload = function() {
    //loadChat()
};