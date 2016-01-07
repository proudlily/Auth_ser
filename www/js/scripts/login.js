var Login = function() {

	var handleLogin = function() {

		$('.login-form').validate({
			errorElement: 'span', //default input error message container
			errorClass: 'help-block', // default input error message class
			focusInvalid: false, // do not focus the last invalid input
			rules: {
				name: {
					required: true
				},
				password: {
					required: true
				},
				remember: {
					required: false
				}
			},

			messages: {
				nick: {
					required: "账户名不能为空哦"
				},
				password: {
					required: "账户密码不能为空哦"
				}
			},

			invalidHandler: function(event, validator) { //display error alert on form submit   
				$('.alert-danger', $('.login-form')).show();
			},

			highlight: function(element) { // hightlight error inputs
				$(element)
					.closest('.form-group').addClass('has-error'); // set error class to the control group
			},

			success: function(label) {
				label.closest('.form-group').removeClass('has-error');
				label.remove();
			},

			errorPlacement: function(error, element) {
				error.insertAfter(element.closest('.input-icon'));
			},

			submitHandler: function(form) {
				return false
			}
		});

		$('.login-form input').keypress(function(e) {
			if (e.which == 13) {
				if ($('.login-form').validate().form()) {
					dologin();
				}
				return false;
			}
		});

		$('.login-form .submit').click(function() {
			if ($('.login-form').validate().form()) {
				dologin();
			}
		});

		var dologin = function() {
			$.ajax({
				data: {
					"u": JSON.stringify(getJson(".login-form"))
				},
				type: "post",
				dataType: "text",
				url: "http://auth.asyou.me.test:9091/Login",
				success: function(data, status) {
					console.log(data);
					console.log(status);
				},
				error: function(data, req) {
					console.log(req);
					alert(req);
				}
			});
		}
	}

	var handleForgetPassword = function() {
		$('.forget-form').validate({
			errorElement: 'span', //default input error message container
			errorClass: 'help-block', // default input error message class
			focusInvalid: false, // do not focus the last invalid input
			ignore: "",
			rules: {
				email: {
					required: true,
					email: true
				}
			},

			messages: {
				email: {
					required: "Email is required."
				}
			},

			invalidHandler: function(event, validator) { //display error alert on form submit   

			},

			highlight: function(element) { // hightlight error inputs
				$(element)
					.closest('.form-group').addClass('has-error'); // set error class to the control group
			},

			success: function(label) {
				label.closest('.form-group').removeClass('has-error');
				label.remove();
			},

			errorPlacement: function(error, element) {
				error.insertAfter(element.closest('.input-icon'));
			},

			submitHandler: function(form) {

			}
		});

		$('.forget-form input').keypress(function(e) {
			if (e.which == 13) {
				if ($('.forget-form').validate().form()) {
					$('.forget-form').submit();
				}
				return false;
			}
		});

		jQuery('#forget-password').click(function() {
			jQuery('.login-form').hide();
			jQuery('.forget-form').show();
		});

		jQuery('#back-btn').click(function() {
			jQuery('.login-form').show();
			jQuery('.forget-form').hide();
		});
	}

	var handleRegister = function() {

		function format(state) {
			if (!state.id) {
				return state.text;
			}
			var $state = $(
				'<span><img src="../assets/global/img/flags/' + state.element.value.toLowerCase() + '.png" class="img-flag" /> ' + state.text + '</span>'
			);

			return $state;
		}

		$('.register-form').validate({
			errorElement: 'span', //default input error message container
			errorClass: 'help-block', // default input error message class
			focusInvalid: false, // do not focus the last invalid input
			ignore: "",
			rules: {
				name: {
					required: true
				},
				email: {
					required: true,
					email: true
				},
				tel: {
					required: true
				},
				nick: {
					required: true
				},
				password: {
					required: true
				},
				rpassword: {
					equalTo: "#register_password"
				},
			},

			messages: {
				nick: {
					required: "请输入昵称哦"
				},
				name: {
					required: "账户名不能为空哦"
				},
				email: {
					required: "邮箱不能为空哦",
					email: "邮箱格式不正确呢"
				},
				tel: {
					required: "请输入手机号码哦"
				},
				password: {
					required: "账户密码不能为空哦"
				},
				rpassword: {
					equalTo: "两次输入的密码不一样哦"
				},
			},

			invalidHandler: function(event, validator) { //display error alert on form submit   

			},

			highlight: function(element) { // hightlight error inputs
				$(element)
					.closest('.form-group').addClass('has-error'); // set error class to the control group
			},

			success: function(label) {
				label.closest('.form-group').removeClass('has-error');
				label.remove();
			},

			errorPlacement: function(error, element) {
				if (element.attr("name") == "tnc") { // insert checkbox errors after the container                  
					error.insertAfter($('#register_tnc_error'));
				} else if (element.closest('.input-icon').size() === 1) {
					error.insertAfter(element.closest('.input-icon'));
				} else {
					error.insertAfter(element);
				}
			},

			submitHandler: function(form) {
				return false
			}
		});

		$('.register-form input').keypress(function(e) {
			if (e.which == 13) {
				if ($('.register-form').validate().form()) {
					doregister();
				}
				return false;
			}
		});

		$('#register-submit-btn').click(function() {
			if ($('.register-form').validate().form()) {
				doregister();
			}
		});

		var doregister = function() {
			$.ajax({
				data: {
					"u": JSON.stringify(getJson(".register-form"))
				},
				type: "post",
				dataType: "text",
				url: "http://auth.asyou.me.test:9091/Register",
				success: function(data, status) {
					$("#register-back-btn").click();
				},
				error: function(data, req) {
					console.log(req);
					alert(req);
				}
			});
		}

		jQuery('#register-btn').click(function() {
			jQuery('.login-form').hide();
			jQuery('.register-form').show();
		});

		jQuery('#register-back-btn').click(function() {
			jQuery('.login-form').show();
			jQuery('.register-form').hide();
		});
	}

	var getJson = function(id) {
		var data_array = $(id).serializeArray();
		var data_return = {};
		for (var i = data_array.length - 1; i >= 0; i--) {
			data_return[data_array[i].name] = data_array[i].value;
		};
		return data_return;
	}

	return {
		//main function to initiate the module
		init: function() {
			handleLogin();
			handleForgetPassword();
			handleRegister();
		}
	};
}();

jQuery(document).ready(function() {
	Login.init();
});