// modules are defined as an array
// [ module function, map of requires ]
//
// map of requires is short require name -> numeric require
//
// anything defined in a previous bundle is accessed via the
// orig method which is the require for previous bundles
parcelRequire = (function (modules, cache, entry, globalName) {
  // Save the require from previous bundle to this closure if any
  var previousRequire = typeof parcelRequire === 'function' && parcelRequire;
  var nodeRequire = typeof require === 'function' && require;

  function newRequire(name, jumped) {
    if (!cache[name]) {
      if (!modules[name]) {
        // if we cannot find the module within our internal map or
        // cache jump to the current global require ie. the last bundle
        // that was added to the page.
        var currentRequire = typeof parcelRequire === 'function' && parcelRequire;
        if (!jumped && currentRequire) {
          return currentRequire(name, true);
        }

        // If there are other bundles on this page the require from the
        // previous one is saved to 'previousRequire'. Repeat this as
        // many times as there are bundles until the module is found or
        // we exhaust the require chain.
        if (previousRequire) {
          return previousRequire(name, true);
        }

        // Try the node require function if it exists.
        if (nodeRequire && typeof name === 'string') {
          return nodeRequire(name);
        }

        var err = new Error('Cannot find module \'' + name + '\'');
        err.code = 'MODULE_NOT_FOUND';
        throw err;
      }

      localRequire.resolve = resolve;
      localRequire.cache = {};

      var module = cache[name] = new newRequire.Module(name);

      modules[name][0].call(module.exports, localRequire, module, module.exports, this);
    }

    return cache[name].exports;

    function localRequire(x){
      return newRequire(localRequire.resolve(x));
    }

    function resolve(x){
      return modules[name][1][x] || x;
    }
  }

  function Module(moduleName) {
    this.id = moduleName;
    this.bundle = newRequire;
    this.exports = {};
  }

  newRequire.isParcelRequire = true;
  newRequire.Module = Module;
  newRequire.modules = modules;
  newRequire.cache = cache;
  newRequire.parent = previousRequire;
  newRequire.register = function (id, exports) {
    modules[id] = [function (require, module) {
      module.exports = exports;
    }, {}];
  };

  var error;
  for (var i = 0; i < entry.length; i++) {
    try {
      newRequire(entry[i]);
    } catch (e) {
      // Save first error but execute all entries
      if (!error) {
        error = e;
      }
    }
  }

  if (entry.length) {
    // Expose entry point to Node, AMD or browser globals
    // Based on https://github.com/ForbesLindesay/umd/blob/master/template.js
    var mainExports = newRequire(entry[entry.length - 1]);

    // CommonJS
    if (typeof exports === "object" && typeof module !== "undefined") {
      module.exports = mainExports;

    // RequireJS
    } else if (typeof define === "function" && define.amd) {
     define(function () {
       return mainExports;
     });

    // <script>
    } else if (globalName) {
      this[globalName] = mainExports;
    }
  }

  // Override the current require with this new one
  parcelRequire = newRequire;

  if (error) {
    // throw error from earlier, _after updating parcelRequire_
    throw error;
  }

  return newRequire;
})({"comment-component.ts":[function(require,module,exports) {
"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.KBCommentComponent = void 0;

var KBCommentComponent = function () {
  function KBCommentComponent(config) {
    var _a;

    this.config = config;
    this.submitting = false;

    this.success = function () {};

    this.config = config;
    this.element = document.createElement("form");
    this.element.className = "kb-comment-form";
    this.element.action = "javascript:";
    this.element.innerHTML = "\n\t\t<div class=\"user-info clearfix\">\n\t\t\t<div class=\"is-reply\">\n\t\t\t\t\u56DE\u590D \u7528\u6237\u540D <a class=\"reset-reply\" href=\"javascript:\">\u53D6\u6D88</a>\n\t\t\t\t<input type=\"hidden\" name=\"parentId\" value=\"0\" />\n\t\t\t\t<input type=\"hidden\" name=\"rId\" value=\"0\" />\n\t\t\t</div>\n\t\t\t<div class=\"input-col\">\n\t\t\t\t<input type=\"text\" name=\"nickname\" maxlength=\"40\" placeholder=\"\u6635\u79F0(\u5FC5\u586B)\" required/>\n\t\t\t</div>\n\t\t\t<div class=\"input-col\">\n\t\t\t\t<input type=\"email\" name=\"mail\" placeholder=\"\u90AE\u7BB1(\u5FC5\u586B)\" required/>\n\t\t\t</div>\n\t\t\t<div class=\"input-col\">\n\t\t\t\t<input type=\"url\" name=\"site\" maxlength=\"40\" placeholder=\"\u7F51\u5740\" />\n\t\t\t</div>\n\t\t</div>\n    <div class=\"message\">\n      <textarea row=\"6\" name=\"content\" placeholder=\"\u8BF7\u8F93\u5165\u4F60\u7684\u7559\u8A00\" required></textarea>\n\t\t</div>\n\t\t<div class=\"btn-group\">\n\t\t\t<button type=\"submit\">\u8BC4\u8BBA</button>\n\t\t</div>";
    this.parentIdField = this.element.querySelector("input[name=parentId]");
    this.rIdField = this.element.querySelector("input[name=rId]");
    this.nickNameField = this.element.querySelector("input[name=nickname]");
    this.mailField = this.element.querySelector("input[name=mail]");
    this.siteField = this.element.querySelector("input[name=site]");
    this.contentField = this.element.querySelector("textarea");
    this.submitBtn = this.element.querySelector("button[type=submit]");
    (_a = this.element.querySelector("a.reset-reply")) === null || _a === void 0 ? void 0 : _a.addEventListener("click", this.resetReply.bind(this));
    this.element.addEventListener("submit", this.onSubmitComment.bind(this), false);
  }

  KBCommentComponent.prototype.setEvent = function (successFunc) {
    this.success = successFunc;
  };

  KBCommentComponent.prototype.getModel = function () {
    return {
      parentId: this.parentIdField.value,
      rId: this.rIdField.value,
      nickName: this.nickNameField.value,
      mail: this.mailField.value,
      site: this.siteField.value,
      content: this.contentField.value,
      ip: returnCitySN.cip,
      token: this.config.token
    };
  };

  KBCommentComponent.prototype.onSubmitComment = function (event) {
    var _this = this;

    event.preventDefault();

    if (this.submitting) {
      return;
    }

    this.submitting = true;
    this.submitBtn.disabled = true;
    axios.post(this.config.apiBase + "/comment", this.getModel(), function (res) {
      if (res.data.code == 200) {
        _this.success();
      }
    });
    return false;
  };

  KBCommentComponent.prototype.resetReply = function () {
    this.setReply(0, 0);
  };

  KBCommentComponent.prototype.setReply = function (parentId, rId) {
    this.rIdField.value = rId.toString();
    this.parentIdField.value = parentId.toString();
  };

  return KBCommentComponent;
}();

exports.KBCommentComponent = KBCommentComponent;
},{}],"kb-comment.ts":[function(require,module,exports) {
"use strict";

var _commentComponent = require("./comment-component");

var KBComment = function () {
  function KBComment() {
    var _this = this;

    this.config = {
      theme: "light",
      apiBase: "",
      token: "light"
    };
    this.container = document.getElementById("kb-comment");

    if (this.container) {
      this.container.className = "kb-comment-container";
      this.config.theme = this.container.dataset.theme || "light";
      this.config.apiBase = this.container.dataset.api || "";
      this.config.token = this.container.dataset.token || location.pathname;
      this.load().then(function () {
        var _a;

        _this.formComponent = new _commentComponent.KBCommentComponent(_this.config);
        (_a = _this.container) === null || _a === void 0 ? void 0 : _a.appendChild(_this.formComponent.element);
      });
    } else {
      console.error("未设定渲染容器");
    }
  }

  KBComment.prototype.load = function () {
    return Promise.all([this.loadTheme(), this.loadLibrary("http://pv.sohu.com/cityjson?ie=utf-8"), this.loadLibrary("https://cdn.jsdelivr.net/npm/axios@0.21.1/dist/axios.min.js"), this.loadLibrary("https://cdn.jsdelivr.net/npm/js-md5@0.7.3/build/md5.min.js")]);
  };

  KBComment.prototype.loadTheme = function () {
    var _this = this;

    return new Promise(function (resolve) {
      var link = document.createElement("link");
      link.rel = "stylesheet";
      link.setAttribute("crossorigin", "anonymous");
      link.onload = resolve;
      link.href = "/themes/" + _this.config.theme + ".css";
      document.head.appendChild(link);
    });
  };

  KBComment.prototype.loadLibrary = function (url) {
    return new Promise(function (resolve) {
      var script = document.createElement("script");
      script.src = url;
      script.onload = resolve;
      document.body.appendChild(script);
    });
  };

  return KBComment;
}();

new KBComment();
},{"./comment-component":"comment-component.ts"}],"../../../../../.nvm/versions/node/v10.16.0/lib/node_modules/parcel-bundler/src/builtins/hmr-runtime.js":[function(require,module,exports) {
var global = arguments[3];
var OVERLAY_ID = '__parcel__error__overlay__';
var OldModule = module.bundle.Module;

function Module(moduleName) {
  OldModule.call(this, moduleName);
  this.hot = {
    data: module.bundle.hotData,
    _acceptCallbacks: [],
    _disposeCallbacks: [],
    accept: function (fn) {
      this._acceptCallbacks.push(fn || function () {});
    },
    dispose: function (fn) {
      this._disposeCallbacks.push(fn);
    }
  };
  module.bundle.hotData = null;
}

module.bundle.Module = Module;
var checkedAssets, assetsToAccept;
var parent = module.bundle.parent;

if ((!parent || !parent.isParcelRequire) && typeof WebSocket !== 'undefined') {
  var hostname = "" || location.hostname;
  var protocol = location.protocol === 'https:' ? 'wss' : 'ws';
  var ws = new WebSocket(protocol + '://' + hostname + ':' + "59247" + '/');

  ws.onmessage = function (event) {
    checkedAssets = {};
    assetsToAccept = [];
    var data = JSON.parse(event.data);

    if (data.type === 'update') {
      var handled = false;
      data.assets.forEach(function (asset) {
        if (!asset.isNew) {
          var didAccept = hmrAcceptCheck(global.parcelRequire, asset.id);

          if (didAccept) {
            handled = true;
          }
        }
      }); // Enable HMR for CSS by default.

      handled = handled || data.assets.every(function (asset) {
        return asset.type === 'css' && asset.generated.js;
      });

      if (handled) {
        console.clear();
        data.assets.forEach(function (asset) {
          hmrApply(global.parcelRequire, asset);
        });
        assetsToAccept.forEach(function (v) {
          hmrAcceptRun(v[0], v[1]);
        });
      } else if (location.reload) {
        // `location` global exists in a web worker context but lacks `.reload()` function.
        location.reload();
      }
    }

    if (data.type === 'reload') {
      ws.close();

      ws.onclose = function () {
        location.reload();
      };
    }

    if (data.type === 'error-resolved') {
      console.log('[parcel] ✨ Error resolved');
      removeErrorOverlay();
    }

    if (data.type === 'error') {
      console.error('[parcel] 🚨  ' + data.error.message + '\n' + data.error.stack);
      removeErrorOverlay();
      var overlay = createErrorOverlay(data);
      document.body.appendChild(overlay);
    }
  };
}

function removeErrorOverlay() {
  var overlay = document.getElementById(OVERLAY_ID);

  if (overlay) {
    overlay.remove();
  }
}

function createErrorOverlay(data) {
  var overlay = document.createElement('div');
  overlay.id = OVERLAY_ID; // html encode message and stack trace

  var message = document.createElement('div');
  var stackTrace = document.createElement('pre');
  message.innerText = data.error.message;
  stackTrace.innerText = data.error.stack;
  overlay.innerHTML = '<div style="background: black; font-size: 16px; color: white; position: fixed; height: 100%; width: 100%; top: 0px; left: 0px; padding: 30px; opacity: 0.85; font-family: Menlo, Consolas, monospace; z-index: 9999;">' + '<span style="background: red; padding: 2px 4px; border-radius: 2px;">ERROR</span>' + '<span style="top: 2px; margin-left: 5px; position: relative;">🚨</span>' + '<div style="font-size: 18px; font-weight: bold; margin-top: 20px;">' + message.innerHTML + '</div>' + '<pre>' + stackTrace.innerHTML + '</pre>' + '</div>';
  return overlay;
}

function getParents(bundle, id) {
  var modules = bundle.modules;

  if (!modules) {
    return [];
  }

  var parents = [];
  var k, d, dep;

  for (k in modules) {
    for (d in modules[k][1]) {
      dep = modules[k][1][d];

      if (dep === id || Array.isArray(dep) && dep[dep.length - 1] === id) {
        parents.push(k);
      }
    }
  }

  if (bundle.parent) {
    parents = parents.concat(getParents(bundle.parent, id));
  }

  return parents;
}

function hmrApply(bundle, asset) {
  var modules = bundle.modules;

  if (!modules) {
    return;
  }

  if (modules[asset.id] || !bundle.parent) {
    var fn = new Function('require', 'module', 'exports', asset.generated.js);
    asset.isNew = !modules[asset.id];
    modules[asset.id] = [fn, asset.deps];
  } else if (bundle.parent) {
    hmrApply(bundle.parent, asset);
  }
}

function hmrAcceptCheck(bundle, id) {
  var modules = bundle.modules;

  if (!modules) {
    return;
  }

  if (!modules[id] && bundle.parent) {
    return hmrAcceptCheck(bundle.parent, id);
  }

  if (checkedAssets[id]) {
    return;
  }

  checkedAssets[id] = true;
  var cached = bundle.cache[id];
  assetsToAccept.push([bundle, id]);

  if (cached && cached.hot && cached.hot._acceptCallbacks.length) {
    return true;
  }

  return getParents(global.parcelRequire, id).some(function (id) {
    return hmrAcceptCheck(global.parcelRequire, id);
  });
}

function hmrAcceptRun(bundle, id) {
  var cached = bundle.cache[id];
  bundle.hotData = {};

  if (cached) {
    cached.hot.data = bundle.hotData;
  }

  if (cached && cached.hot && cached.hot._disposeCallbacks.length) {
    cached.hot._disposeCallbacks.forEach(function (cb) {
      cb(bundle.hotData);
    });
  }

  delete bundle.cache[id];
  bundle(id);
  cached = bundle.cache[id];

  if (cached && cached.hot && cached.hot._acceptCallbacks.length) {
    cached.hot._acceptCallbacks.forEach(function (cb) {
      cb();
    });

    return true;
  }
}
},{}]},{},["../../../../../.nvm/versions/node/v10.16.0/lib/node_modules/parcel-bundler/src/builtins/hmr-runtime.js","kb-comment.ts"], null)
//# sourceMappingURL=/kb-comment.8d40b601.js.map