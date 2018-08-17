/*
Sniperkit-Bot
- Status: analyzed
*/

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
            var connection = new WebSocket('ws:' + location.host + '/' + this.data.integrationName + '/cmd');
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
            return ansi_up.ansi_to_html(ansi_up.escape_for_html(value));
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
            if (this.data.executor != null) {
                this.data.executor.isSelected = true;
                this.$parent.showExecutor(this.data.executor);
            }
        },
        run: function () {
            var newExecutor = {
                name: this.data.TaskName,
                taskName: this.data.TaskName,
                integrationName: this.data.IntegrationName,
                combinedOutput: "",
                status: "",
                error: "",
                isSelected: false,
            };
            this.data.executor = newExecutor;
            this.data.executor.isSelected = true;
            this.$parent.startExecutor(newExecutor);
            this.$parent.showExecutor(this.data.executor);
        }
    }
});

new Vue({
    el: '#taskrunner',
    data: {
        tasks: [],
        integrations: {},
        executors: [],
        focusedExecutor: null,
        textFilter: "",
    },
    created: function () {
        var self = this;
        var connection = new WebSocket('ws:' + location.host + '/tasklist');
        connection.onerror = function (error) {
            console.log('WebSocket Error ' + error);
        };

        connection.onmessage = function (e) {
            var taskInfo = JSON.parse(e.data);
            taskInfo.executor = null;
            self.$data.tasks.push(taskInfo);
            Vue.set(self.$data.integrations, taskInfo.IntegrationName, {
                "name": taskInfo.IntegrationName,
                "selected": true
            });
        };

        connection.onclose = function (e) {
            console.log("ALL TASKS LOAD");
        };
    },
    computed: {
        filteredTasks: function () {
            var self = this;
            var filteredTasks = this.$data.tasks.filter(function (task) {
                var integrationIsSelected = self.$data.integrations[task.IntegrationName].selected;
                var taskNameContainsFilter = task.TaskName.indexOf(self.$data.textFilter) > -1;
                return integrationIsSelected && taskNameContainsFilter;
            });
            filteredTasks.sort(function (a, b) {
                if (a.TaskName < b.TaskName) return -1;
                if (a.TaskName > b.TaskName) return 1;
                return 0;
            });
            return filteredTasks;
        }
    },
    methods: {
        toggleIntegrationFilter: function (integration) {
            integration.selected = !integration.selected;
        },
        showExecutor: function (executor) {
            for (var k in this.$data.executors) {
                var currentExecutor = this.$data.executors[k];
                if (executor != currentExecutor) {
                    currentExecutor.isSelected = false;
                }
            }
            this.focusedExecutor = executor;
        },
        startExecutor: function (executor) {
            console.log(executor);

            this.executors.push(executor);
        }
    }
});