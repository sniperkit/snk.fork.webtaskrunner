
Vue.component('taskrunner',{
  props:['data'],
  template:'#taskrunner',
  methods:{
      navigateTo:function(targetPath){
          location.href = targetPath;
      }
  },
})

new Vue({
    el:'#overview',
    data:{
        taskrunners:[]
    },
    created:function(){
    this.$http.get('/taskrunners')
        .then(function(response){
            if(response.ok){
                var tasks = response.body;
                for(k in tasks){
                    var newTaskrunner={
                        name:tasks[k].Name,
                        imageUrl:tasks[k].ImageUrl,
                        route:tasks[k].Route,
                    };
                    console.log(newTaskrunner);
                    this.taskrunners.push(newTaskrunner);
                }
            }else{
                console.log("error while receiving task list");
            }
        });
    },
});