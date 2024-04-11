document.addEventListener("DOMContentLoaded",function (){
    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');
    const role = urlParams.get('role')
    console.log(productId)
    fetch(`http://localhost:9090/viewEach/${productId}`)
.then(response => {
    if (!response.ok){
        throw new Error('Network response was not ok');
    }
   return response.json()
})
.then(data => {
    
    data.forEach(element => {
       viewCard(element);
    });
})
.catch(error => console.error('Error fetching product data:', error))  
})
function viewCard(el){
const container = document.getElementById("dCard")
const urlParams = new URLSearchParams(window.location.search);
const role = urlParams.get('role');
const id = urlParams.get('id')
const card = document.createElement("div")
card.classList.add("card","mb-3")
let htmlcontent = `
            <div class="card-body">
            <div class="card-body d-flex justify-content-center align-items-center">
              <img src="${el.path}" class="card-img-top" alt="..." style="width: 100px; height: 200px;">
            </div>

            <div class="card-body d-flex justify-content-center align-items-center">
              <h5 class="card-title">${el.mName}</h5>                     
            </div>
            <div class="card-body d-flex justify-content-center align-items-center">
              <p>PRICE : ${el.mPrice}</p>                    
            </div>
            <div class="card-body d-flex justify-content-center align-items-center">
              <p>SPECS : ${el.mSpecs}</p>                      
            </div> 


            
`;
if(role === 'admin') {
    htmlcontent += `
    <div class="card-body d-flex justify-content-center align-items-center">
        <button class="btn btn-danger mr-2"  onclick="deleteItem('${el.mid}')">Delete</button>
        <button class="btn btn-warning mx-2" onclick="editItem('${el.mid}')">Edit</button>


        </div>
    `;
}
htmlcontent += `</html>`           
card.innerHTML = htmlcontent
container.appendChild(card)
}
function editItem(id){
console.log(id)
window.location.href = `editpage.html?id=${id}`
}
function deleteItem(id){
const jsonData = JSON.stringify({
    mid: parseInt(id) 
});

fetch("http://localhost:9090/delete",{
    method : "DELETE",
    body: jsonData
     
})
  .then(response => {
    if (!response.ok) {
    throw new Error("Network response was not ok");
    }
    return response.json();
})
  .then(data => {
    alert("Deleted successfully")
    window.location.href = "admin.html"
    console.log(data);
  })
  .catch(error => {
    console.error("Error during request:", error);
    })
}