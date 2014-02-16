
var Task = function(json) {
    
    var name = json.name;
    var description = json.description;
    console.log(name);  
    console.log(description);

    

}

var TaskApI = new function() {
    var thisCache = this;    

    this.createNewTask = function(params) {

        thisCache.taskResult = [];

        var data = {                      
            name:         "Jay Kan",
            description:  "New Task"                                                
        }
        

        var domain = "localhost:6060";
        var postUrl = "http://" + domain + "/api/tasks";

        $.ajax({
            url:  postUrl,
            data: JSON.stringify(data),
            type: 'POST',
            dataType: 'json',
            success: function(jsonResponse) {                                  
                console.log(jsonResponse);
                var newTask = new Task(jsonResponse);
                console.log(newTask);                   
            },
            error: function() { } 
        });
    };
}


