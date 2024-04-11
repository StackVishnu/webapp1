let productId
document.addEventListener("DOMContentLoaded", function (){
    const urlParams = new URLSearchParams(window.location.search);
    productId = urlParams.get('id');
    const role = urlParams.get('role');
    console.log(productId);
    
    fetch(`http://localhost:9090/viewEach/${productId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.length > 0) { // Check if data is not empty
                const product = data[0]; // Assuming there's only one product returned
                console.log(product);
                document.getElementById('phone_name').value = product.mName;
                document.getElementById('specs').value = product.mSpecs;
                document.getElementById('price').value = product.mPrice;
            } else {
                console.error('No data found');
            }
        })
        .catch(error => console.error('Error fetching product data:', error));
});
function edit(){
    const form = document.getElementById("add-form");
        const formData = new FormData(form);
        const mid = parseInt(productId)
        formData.append('mid', productId)
        console.log("test"+formData)
        event.preventDefault()
    fetch(`http://localhost:9090/edit`,{
        method : "PUT",
        body : formData
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.json();
    })
    .then(data => {
        alert("Product edited successfully")
        window.location.href = "admin.html"
        console.log(data);
    })
    .catch(error => {
        console.error("Error during form submission:", error);
    });
}


