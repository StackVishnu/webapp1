document.getElementById("submitBtn").addEventListener("click", function() {
    const form = document.getElementById("add-form");
    const formData = new FormData(form);
    event.preventDefault()

fetch("http://localhost:9090/insert", {
    method: "POST",
    body: formData
})
.then(response => {
    if (!response.ok) {
        throw new Error("Network response was not ok");
    }
    return response.json();
})
.then(data => {
    console.log(data);
    alert("Product added successfully")
    window.location.href ="admin.html"
})
.catch(error => {
    console.error("Error during form submission:", error);
});
});