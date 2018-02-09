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
    
    
    
    function loginForm() {
    
        var id,pass;
        
        id=document.getElementById("phoneinpt").value;
        pass=document.getElementById("passinpt").value;
        
    
        if(!pass && !id){
            alert("Fill the forms carefully...");
            location.reload();
        }else{
           
            var data="id="+id+"&name=user"+"&pass="+pass+"&add=esf#123433#765432#";
            alert(data);
            var xhttp = new XMLHttpRequest();
            xhttp.onreadystatechange = function() {
              if (this.readyState == 4 && this.status == 200) {
                
                if (typeof(Storage) !== "undefined") {
                    // Store
                    
                    localStorage.setItem("magic-login", this.responseText);
                    // Retrieve
                    alert(localStorage.getItem("magic-login"));
                }
              }
            };
    
           // for (i=0;i<2;i++){ 
            xhttp.open("POST", "http://localhost:3000/login", true);
            xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xhttp.send(data);
          //}
        }
       
      }
      
      
      
    
    