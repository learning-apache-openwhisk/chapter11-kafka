// main.js 

var base = location.search
if(base == "") {
	base = location.href
	base = base.substring(0, base.lastIndexOf("/"))
	base = base.substring(0, base.lastIndexOf("/"))
} else {
	base = base.substring(1)
}
console.log(base)
var sender = base + "/sender"
var receiver = base + "/receiver"
var nickname = ""
var group = ""

// poll the chat
function poll() {
	$.post(receiver, {"group": group},  
		function (data) {
			//console.log(data)
			if(data.messages.lenght==0)
				return
			var curr = $("#room").val()
			for(message of data.messages) {
				curr += message + "\n"
			}
			$room = $("#room")
			$room.val(curr)
			$room.scrollTop($room[0].scrollHeight)
		})
}

// join the chat
function join() {
	nick = $("#nick").val()
	pass = $("#pass").val()
	if(nick == "" || pass== "") {
		alert("Please specify nickname and password")
		return
	}
	// first connection, get the group
	$.post(receiver, {"nick": nick, "pass": pass},
		function(data) {
			//console.log(data)
			if(data.error) {
				alert(data.error)
				return
			}
			// logged in, initialize
			if(data.group) { 
				nickname = nick
				group = data.group
				$("#form").hide()
				$("#me").text(nick)
				$("#message").removeAttr("disabled")
				$.post(sender, {"message": "**** "+nick +" joined ****"})
				setInterval(poll, 3000)
			}
		})
}

function message(e) {
	message = $("#message").val()
	if(message == "") 
		return
	if(e.keyCode === 13) { 
		//alert(message)
		$("#message").val("")
		$.post(sender, 
			{"message": "["+nickname+"] "+message}, 
			function (data) {
			  console.log(data)
			}
		)
	}
}

$(function() {
	$("#join").click(join)
	$("#message").keyup(message)
})
