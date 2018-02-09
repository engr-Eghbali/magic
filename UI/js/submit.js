$(document).ready(function(){


var elementPosition = $('#header').offset();

$(window).scroll(function(){
        if($(window).scrollTop() > elementPosition.top){
            $('#header').css('position','fixed').css('top','0');
              $('#menu').css('opacity','0.5');
        } else {
            $('#header').css('position','static');
            $('#menu').css('opacity','0.2');
        }    
});

});



function submitForm() {

    var name,id,pass1,pass2,add;
    name=document.getElementById("nameinpt").value;
    id=document.getElementById("phoneinpt").value;
    pass1=document.getElementById("passinpt").value;
    pass2=document.getElementById("passinpt2").value;

    if(pass1!=pass2 || !id){
        alert("Fill the forms carefully...");
        location.reload();
    }else{

        var data="id="+id+"&name="+name+"&pass="+pass1+"&add=ESF#123433#765432#";
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
          if (this.readyState == 4 && this.status == 200) {
            
            if (typeof(Storage) !== "undefined") {
                // Store
                
                localStorage.setItem("magic-sub", this.responseText);
                // Retrieve
                alert(localStorage.getItem("magic-sub"));
            }
          }
        };

       // for (i=0;i<2;i++){ 
        xhttp.open("POST", "http://localhost:3000/submit", true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send(data);
      //}
    }
   
  }
  
  
  

