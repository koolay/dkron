var dkron = angular.module('dkron', ['angular-rickshaw']);

dkron.filter('statusClass', function() {
  var friendlyStatusFilter = function(job) {
    if (job.disabled) {
      return 'text-muted glyphicon-ban-circle'
    }
    switch(job.status) {
      case 'success':
        return 'status-success glyphicon-ok-sign'
      case 'failed':
        return 'status-danger glyphicon-remove-sign'
      case 'partially_failed':
        return 'status-warning glyphicon-exclamation-sign'
      case 'running':
        return 'status-running glyphicon-play-circle'
      }
    return input;
  };
  return friendlyStatusFilter;
});

dkron.constant('hideDelay', 2000);

dkron.controller('JobListCtrl', function ($scope, $http, $interval, hideDelay) {
  $scope.runJob = function(jobName) {
    $scope["running_" + jobName] = true;
    var response = $http.post(DKRON_API_PATH + '/jobs/' + jobName);
    response.success(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-success fade in">Success running job ' + jobName + '</div>');

      $(".alert-success").delay(hideDelay).slideUp(200, function(){
        $(".alert").alert('close');
        window.location.reload();
      });
    });

    response.error(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-danger fade in"><button type="button" class="close close-alert" data-dismiss="alert" aria-hidden="true">x</button>Error running job ' + jobName + '</div>');
    });
  };

  $scope.deleteJob = function(jobName) {
    $scope["deleting_" + jobName] = true;
    var response = $http.delete(DKRON_API_PATH + '/jobs/' + jobName);
    response.success(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-success fade in">Successfully removed job ' + jobName + '</div>');

      $(".alert-success").delay(hideDelay).slideUp(200, function(){
        $(".alert").alert('close');
        window.location.reload();
      });
    });

    response.error(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-danger fade in"><button type="button" class="close close-alert" data-dismiss="alert" aria-hidden="true">x</button>Error removing job ' + jobName + '</div>');
    });
  };

  $scope.toggleJob = function(jobName) {
    var response = $http.post(DKRON_API_PATH + '/jobs/' + jobName + '/toggle');
    response.success(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-success fade in">Successfully toggled job ' + jobName + '</div>');

      $(".alert-success").delay(hideDelay).slideUp(200, function(){
        $(".alert").alert('close');
        window.location.reload();
      });
    });

    response.error(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-danger fade in"><button type="button" class="close close-alert" data-dismiss="alert" aria-hidden="true">x</button>Error toggle job ' + jobName + '</div>');
    });
  }

  var updateView = function() {
    var response = $http.get(DKRON_API_PATH + '/jobs');
    response.success(function(data, status, headers, config) {
      $scope.updateStatus(data);

      $("#conn-error").delay(hideDelay).slideUp(200, function(){
        $("#conn-error").alert('close');
      });
    });

    response.error(function(data, status, headers, config) {
      $('#message').html('<div id="conn-error" class="alert alert-danger fade in"><button type="button" class="close close-alert" data-dismiss="alert" aria-hidden="true">x</button>Error getting data</div>');
    });
  }

  // calculate page in place
  $scope.groupToPages = function () {
    $scope.gap = Math.round($scope.jobs.length / $scope.itemsPerPage);
    $scope.pagedItems = [];
    
    for (var i = 0; i < $scope.jobs.length; i++) {
        if (i % $scope.itemsPerPage === 0) {
            $scope.pagedItems[Math.floor(i / $scope.itemsPerPage)] = [ $scope.jobs[i] ];
        } else {
            $scope.pagedItems[Math.floor(i / $scope.itemsPerPage)].push($scope.jobs[i]);
        }
    }
  };

  $scope.range = function(size, start, end) {
      var ret = [];

      if (size < end) {
          end = size;
          start = size-$scope.gap;
      }
      for (var i = start; i < end; i++) {
          ret.push(i);
      }
      return ret;
  };

  $scope.prevPage = function() {
      if ($scope.currentPage > 0) {
          $scope.currentPage--;
      }
  };

  $scope.nextPage = function() {
      if ($scope.currentPage < $scope.pagedItems.length - 1) {
          $scope.currentPage++;
      }
  };

  $scope.setPage = function() {
      $scope.currentPage = this.n;
  };

  // init
  $scope.gap = 0;
  $scope.groupedItems = [];
  $scope.itemsPerPage = 15;
  $scope.pagedItems = [];
  $scope.currentPage = 0;
  
  $scope.success_count = 0;
  $scope.error_count = 0;
  $scope.failed_jobs = 0;
  $scope.successful_jobs = 0;
  $scope.jobs = [];

  $scope.updateStatus = function(data) {
    var success_count = 0;
    var error_count = 0;
    $scope.jobs = data;

    $scope.failed_jobs = 0;
    $scope.successful_jobs = 0;

    // functions have been describe process the data for display
    $scope.groupToPages();

    for(i=0; i < data.length; i++) {
      success_count = success_count + data[i].success_count;
      error_count = error_count + data[i].error_count;

      if (new Date(Date.parse(data[i].last_success)) > new Date(Date.parse(data[i].last_error))) {
        $scope.successful_jobs = $scope.successful_jobs + 1;
      } else {
        $scope.failed_jobs = $scope.failed_jobs + 1;
      }
    }

    $scope.success_count = success_count;
    $scope.error_count = error_count;
  }

  updateView();
  hljs.initHighlightingOnLoad();

});

dkron.controller('IndexCtrl', function ($scope, $http, $interval, $element) {
  $scope.options = {
    renderer: 'line',
    interpolation: 'linear'
  };

  $scope.series = [{
      name: 'Success count',
      color: 'green',
      data: [{x: 0, y: 0}]
  },{
      name: 'Error count',
      color: 'red',
      data: [{x: 0, y: 0}]
  }];
  $scope.features = {
      hover: {
          xFormatter: function(x) {
              return x;
          },
          yFormatter: function(y) {
              return y;
          }
      },
      legend: {
        toggle: false,
        highlight: true
      },
      yAxis: {
        tickFormat: Rickshaw.Fixtures.Number.formatKMBT
      },
      xAxis: {
        tickFormat: Rickshaw.Fixtures.Number.formatKMBT,
        timeUnit: 'hour'
      }
  };

  updateView = function() {
    var response = $http.get(DKRON_API_PATH + '/jobs');
    response.success(function(data, status, headers, config) {
      $scope.updateGraph(data);
    });

    response.error(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-danger fade in"><button type="button" class="close close-alert" data-dismiss="alert" aria-hidden="true">x</button>Error getting data</div>');
    });

    var mq = $http.get(DKRON_API_PATH + '/members');
    mq.success(function(data, status, headers, config) {
      angular.forEach(data, function(val, key){
          switch(val.Status) {
            case 0:
              data[key].Status = "none";
              break;
            case 1:
              data[key].Status = "alive";
              break;
            case 2:
              data[key].Status = "leaving";
              break;
            case 3:
              data[key].Status = "left";
              break;
            case 4:
              data[key].Status = "failed";
              break;
          }
      });
      $scope.members = data;
    });

    mq.error(function(data, status, headers, config) {
      $('#message').html('<div class="alert alert-danger fade in"><button type="button" class="close close-alert" data-dismiss="alert" aria-hidden="true">x</button>Error getting data</div>');
    });
  }

  $scope.success_count = 0;
  $scope.error_count = 0;
  $scope.failed_jobs = 0;
  $scope.successful_jobs = 0;
  $scope.jobs = [];

  $scope.updateGraph = function(data) {
    var gdata = $scope.series[0].data;
    var name = $scope.series[0].name;
    var color = $scope.series[0].color;
    var success_count = 0;
    var error_count = 0;
    var diff = 0;

    $scope.jobs = data;
    $scope.failed_jobs = 0;
    $scope.successful_jobs = 0;

    for(i=0; i < data.length; i++) {
      success_count = success_count + data[i].success_count;
      error_count = error_count + data[i].error_count;

      if (new Date(Date.parse(data[i].last_success)) > new Date(Date.parse(data[i].last_error))) {
        $scope.successful_jobs = $scope.successful_jobs + 1;
      } else {
        $scope.failed_jobs = $scope.failed_jobs + 1;
      }
    }
    if($scope.success_count != 0) {
      diff = success_count - $scope.success_count;
    }
    $scope.success_count = success_count;

    gdata.push({x: gdata[gdata.length - 1].x + 1, y: diff})

    $scope.series[0] = {
      name: name,
      color: color,
      data: gdata
    };

    gdata = $scope.series[1].data;
    name = $scope.series[1].name;
    color = $scope.series[1].color;

    if($scope.error_count != 0) {
      diff = error_count - $scope.error_count;
    }
    $scope.error_count = error_count;

    gdata.push({x: gdata[gdata.length - 1].x + 1, y: diff})

    $scope.series[1] = {
      name: name,
      color: color,
      data: gdata
    };
  }

  $interval(function() {
    updateView();
  }, 2000);

  updateView();
});
