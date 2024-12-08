window.onload = function () {
    var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("form").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        conn.send(`{"id":1,"chatId":1,"userId":1,"text":"${msg.value}"}`);
        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://localhost:8090/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                if(messages[i] != "") {
                    console.log("message: " + messages[i])
                    var obj = JSON.parse(messages[i])
                    console.log("obj: " + obj.text)
                    var item = document.createElement("div");
                    item.innerText = obj.text;
                    appendLog(item);
                } 
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};