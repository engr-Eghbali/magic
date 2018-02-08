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
    
    