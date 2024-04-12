let productId
document.addEventListener("DOMContentLoaded", function (){
    const urlParams = new URLSearchParams(window.location.search);
    productId = urlParams.get('id');
    const role = urlParams.get('role');

    
    fetch(`http://localhost:9090/viewEach/${productId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.length > 0) {
                const product = data[0]; 

                document.getElementById('phone_name').value = product.mName;
                document.getElementById('specs').value = product.mSpecs;
                document.getElementById('price').value = product.mPrice;
            } else {
                console.error('No data found');
            }
        })
        .catch(error => alert("alert"));
});
function updateProduct(event) {
    event.preventDefault();

    // Obtain all form data
    const phoneName = document.getElementById('phone_name').value;
    const specs = document.getElementById('specs').value;
    const price = document.getElementById('price').value;
    const imageFile = document.getElementById('image').files[0];

    // Create FormData object
    const formData = new FormData();
    formData.append('phone_name', phoneName);
    formData.append('specs', specs);
    formData.append('price', price);
    formData.append('image', imageFile);
    formData.append('mid', productId);

    // Assuming the API endpoint for updating the product is '/edit'
    fetch('http://localhost:9090/edit', {
        method: 'POST',
        body: formData // Send FormData for form data upload
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        // Redirect to home page after successful update
        window.location.href = 'admin.html'; // Update with your home page URL
    })
    .catch(error => {
        console.error('Error updating product:', error);
        console.log("Error name:", error.name);
            console.log("Error message:", error.message);
            console.log("Error stack trace:", error.stack);
            alert("Error editing product. Please try again later.");
        // Handle error if needed
    });
}
