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
            alert("Fill the forms");
            location.reload();
        }else{
    
            var data="id="+id+"&name=login"+"&pass="+pass+"&add=1234.33#7654.32";
            var xhttp = new XMLHttpRequest();
            xhttp.onreadystatechange = function() {
              if (this.readyState == 4 && this.status == 200) {
                alert("ready and status 200");
              }else{
                  alert("check your connection");
              }
            };
           // for (i=0;i<2;i++){
            xhttp.open("POST", "http://localhost:3000/login", true);
            xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xhttp.send(data);
          //}
               var  server_response= responseText;
          
               if (typeof(Storage) !== "undefined") {
                // Store
                
                localStorage.setItem("magic-sub", server_response);
                // Retrieve
               alert(localStorage.getItem("magic"));
            } else {
                alert("cant support webstorage...");
            }
          
        }
       
      }
      
      
      
    
    