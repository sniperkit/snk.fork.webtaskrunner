
Vue.component('executor', {
    props: ['data'],
    template: '#executor',
    created: function () {
        this.run();
    },
    methods: {
        run: function () {
            this.data.status = "running";
            var self = this;
            var connection = new WebSocket('ws:' + location.host + location.pathname + '/cmd', ['chat']);
            // When the connection is open, send some data to the server
            connection.onopen = function () {
                connection.send(self.data.taskName);
            };

            connection.onerror = function (error) {
                self.data.combinedOutput += error;
                console.log('WebSocket Error ' + error);
            };

            connection.onmessage = function (e) {
                response = JSON.parse(e.data);
                if (response.Status == 1) {
                    self.data.combinedOutput += response.Line;
                } else if (response.Status == 2) {
                    self.data.error = response.Error;
                }
                console.log(response);
            };

            connection.onclose = function (e) {
                self.data.status = "done";
            };
        },
        ansiUp: function (value) {
            return ansi_up.ansi_to_html(value);
        },
        htmlCarriageReturn: function (value) {
            return value.replace(new RegExp('\n', 'g'), "<br>");
        }
    }
});

Vue.component('task', {
    props: ['data'],
    template: '#task',
    created: function () {
    },
    methods: {
        showExecutor: function () {
            this.$parent.showExecutor(this.data.executor);
        },
        run: function () {
            var newExecutor = {
                name: this.data.name,
                taskName: this.data.name,
                combinedOutput: "",
                status: "",
                error: "",
            };
            this.data.executor = newExecutor;
            this.$parent.startExecutor(newExecutor);
            this.$parent.showExecutor(this.data.executor);
        }
    }
});

new Vue({
    el: '#taskrunner',
    data: {
        tasks: [],
        executors: [],
        focusedExecutor: null,
    },
    created: function () {
        this.$http.get(location.pathname + '/tasks')
            .then(function (response) {
                if (response.ok) {
                    var tasks = response.body;
                    for (k in tasks) {
                        var newTask = {
                            name: tasks[k],
                            executor: null,
                        };
                        this.tasks.push(newTask);
                    }
                } else {
                    console.log("error while receiving task list");
                }
            });
    },
    methods: {
        showExecutor: function (executor) {
            this.focusedExecutor = executor;
        },
        startExecutor: function (executor) {
            console.log(executor);

            this.executors.push(executor);
        }
    }
});