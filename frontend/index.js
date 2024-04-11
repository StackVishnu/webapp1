const loginForm = document.getElementById("login-form") 
loginForm.addEventListener("submit", function(event) {
event.preventDefault(); 

const formData = new FormData(loginForm); 
 


fetch("http://localhost:9090/login", {
    method: "POST",
    body: formData
})
.then(response => {
    if (!response.ok) {
        alert("Wrong Credentials")
        throw new Error("Network response was not ok");
    }
    return response.json(); 
})
.then(data => {
    console.log(data)
     openPage(data)
   
})
.catch(error => {
    
    console.error("Error during login:", error);
    
});
});
function openPage(data){
if(data.role == "admin"){
    window.location.href = "admin.html"
}else{
    window.location.href = "user.html"
}
}