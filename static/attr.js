const title = document.getElementById("title")
const content = document.getElementById("content")
const author = document.getElementById("author")
async function FetchData(id, to_paste){
    fetch("/create", {
        method: "POST",
        headers: {
            'Content-Type': "application/json",
        },
        body: JSON.stringify({
            title: title.value,
            content: content.value,
            author: author.value,
            topicid: parseInt(id),
        })
    })
    .then(res => {
        if(res !== undefined){
            if (res.status === 500) {
                console.log("Server Error")
                return
            } else if (res.status === 400) {
                console.log("Bad Request")
                return
            }

            return res.json()
        }else {
            location.reload()
        }
    })
    .then(data => {
        if (to_paste) {
            if(data!==undefined){
				title.value = ""
				content.value = ""
				author.value = ""
                if(parseInt(data)){
                    window.location = `/paste/${data}`
                    return
                }else
                    location.reload()
            }else{
                location.reload()    
                return
            }
        } else {
            location.reload()
            return
        }
    })
    .catch(err => alert(err))
}
