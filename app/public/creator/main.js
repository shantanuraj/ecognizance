var ArticlesInput = React.createClass({displayName: "ArticlesInput",
	render:function() {
		return (
			React.createElement("div", null, 
				React.createElement("div", {className: "form-group"}, 
				  React.createElement("input", {type: "text", className: "form-control", placeholder: "Enter title", name: "Articles." + this.props.data + ".Title"})
				), 
				React.createElement("div", {className: "form-group"}, 
				  React.createElement("input", {type: "text", className: "form-control", placeholder: "Enter author name", name: "Articles." + this.props.data + ".Author"})
				), 
				React.createElement("div", {className: "form-group"}, 
				  React.createElement("input", {type: "email", className: "form-control", placeholder: "Enter author email", name: "Articles." + this.props.data + ".Email"})
				), 
				React.createElement("div", {className: "form-group"}, 
				  React.createElement("textarea", {className: "form-control", rows: "11", placeholder: "Enter post", name: "Articles." + this.props.data + ".Post"})
				)
			)
		);
	}
});

var EcohustleInput = React.createClass({displayName: "EcohustleInput",
	render: function() {
		return (
			React.createElement("div", null, 
				React.createElement("div", {className: "form-group"}, 
	            	React.createElement("input", {type: "text", className: "form-control", placeholder: "Enter image URL", name: "Ecohustle." + this.props.data + ".Image"})
	          	), 
	          	React.createElement("div", {className: "form-group"}, 
	            	React.createElement("input", {type: "text", className: "form-control", placeholder: "Enter image caption", name: "Ecohustle." + this.props.data + ".Caption"})
	          	)
          	)
		);
	}
});

var ArtcilesInputGroup = React.createClass({displayName: "ArtcilesInputGroup",
	render:function() {
		var rows = [];
		for (var i = 0; i < this.props.data; i++) {
		    rows.push(React.createElement(ArticlesInput, {data: i, key: i}));
		}
		return React.createElement("div", null, rows);
	}
})

var EcohustleInputGroup = React.createClass({displayName: "EcohustleInputGroup",
	render:function() {
		var rows = [];
		for (var i = 0; i < this.props.data; i++) {
		    rows.push(React.createElement(EcohustleInput, {data: i, key: i}));
		}
		return React.createElement("div", null, rows);
	}
})

function updateSupplementaryGroup() {
	var i = document.getElementsByClassName('supplementary-selector')[0].value;
	React.render(
		React.createElement(ArtcilesInputGroup, {data: i}),
		document.getElementById('supplementary-input')
	);
}

function updateEcohustleGroup() {
	var i = document.getElementsByClassName('ecohustle-selector')[0].value;
	React.render(
		React.createElement(EcohustleInputGroup, {data: i}),
		document.getElementById('ecohustle-input')
	);
}

