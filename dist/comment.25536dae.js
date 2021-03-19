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

var __assign = void 0 && (void 0).__assign || function () {
  __assign = Object.assign || function (t) {
    for (var s, i = 1, n = arguments.length; i < n; i++) {
      s = arguments[i];

      for (var p in s) {
        if (Object.prototype.hasOwnProperty.call(s, p)) t[p] = s[p];
      }
    }

    return t;
  };

  return __assign.apply(this, arguments);
};

var STORAGE_NAME = "kb-comment-user";

var KBCommentComponent = function () {
  function KBCommentComponent(config) {
    var _a;

    this.config = config;
    this.submitting = false;
    this.config = config;
    this.element = document.createElement("form");
    this.element.className = "kb-comment-form";
    this.element.action = "javascript:";
    this.element.innerHTML = "\n      <input type=\"hidden\" name=\"parentId\" value=\"0\" />\n      <input type=\"hidden\" name=\"rId\" value=\"0\" />\n      <div class=\"user-info clearfix\">\n        <div class=\"input-col\">\n          <input type=\"text\" name=\"nickname\" maxlength=\"40\" placeholder=\"\u6635\u79F0(\u5FC5\u586B)\" required/>\n        </div>\n        <div class=\"input-col\">\n          <input type=\"email\" name=\"mail\" placeholder=\"\u90AE\u7BB1(\u5FC5\u586B)\" required/>\n        </div>\n        <div class=\"input-col\">\n          <input type=\"url\" name=\"site\" maxlength=\"40\" placeholder=\"\u7F51\u5740\" />\n        </div>\n      </div>\n      <div class=\"message\">\n        <textarea row=\"6\" name=\"content\" placeholder=\"\u8BF7\u8F93\u5165\u4F60\u7684\u7559\u8A00\" required></textarea>\n      </div>\n      <div class=\"btn-group\">\n        <button type=\"submit\">\u8BC4\u8BBA</button>\n      </div>";
    this.parentIdField = this.element.querySelector("input[name=parentId]");
    this.rIdField = this.element.querySelector("input[name=rId]");
    this.nickNameField = this.element.querySelector("input[name=nickname]");
    this.mailField = this.element.querySelector("input[name=mail]");
    this.siteField = this.element.querySelector("input[name=site]");
    this.contentField = this.element.querySelector("textarea");
    this.submitBtn = this.element.querySelector("button[type=submit]");
    (_a = this.element.querySelector("a.reset-reply")) === null || _a === void 0 ? void 0 : _a.addEventListener("click", this.resetReply.bind(this));
    this.element.addEventListener("submit", this.onSubmitComment.bind(this), false);
    var info = JSON.parse(window.localStorage.getItem(STORAGE_NAME) || "null");

    if (info) {
      this.nickNameField.value = info.nickName;
      this.mailField.value = info.mail;
      this.siteField.value = info.site;
    }
  }

  KBCommentComponent.prototype.setEvent = function (successFunc) {
    this.success = successFunc;
  };

  KBCommentComponent.prototype.getModel = function () {
    var info = {
      nickName: this.nickNameField.value,
      mail: this.mailField.value,
      site: this.siteField.value
    };
    window.localStorage.setItem(STORAGE_NAME, JSON.stringify(info));
    return __assign(__assign({}, info), {
      parentId: Number(this.parentIdField.value),
      rId: Number(this.rIdField.value),
      content: this.contentField.value,
      articleToken: this.config.token,
      pageUrl: top.location.href,
      pageTitle: top.document.title
    });
  };

  KBCommentComponent.prototype.onSubmitComment = function (event) {
    var _this = this;

    event.preventDefault();

    if (this.submitting) {
      return;
    }

    this.submitting = true;
    this.submitBtn.disabled = true;
    axios.post("api/comment", this.getModel()).then(function (res) {
      _this.submitBtn.disabled = false;
      _this.submitting = false;

      if (res.data.code == 200) {
        _this.success && _this.success();
        _this.contentField.value = "";
      }
    });
    return false;
  };

  KBCommentComponent.prototype.resetReply = function () {
    this.setReply("0", "0");
  };

  KBCommentComponent.prototype.setReply = function (parentId, rId) {
    this.rIdField.value = rId;
    this.parentIdField.value = parentId;
  };

  return KBCommentComponent;
}();

exports.KBCommentComponent = KBCommentComponent;
},{}],"time-ago.ts":[function(require,module,exports) {
"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.timeAgo = timeAgo;
var thresholds = [1000, "秒", 1000 * 60, "分", 1000 * 60 * 60, "时", 1000 * 60 * 60 * 24, "天", 1000 * 60 * 60 * 24 * 7, "周", 1000 * 60 * 60 * 24 * 27, "月"];
var formatOptions = {
  month: "short",
  day: "numeric",
  year: "numeric"
};

function timeAgo(value) {
  var date = new Date(value);
  var elapsed = new Date().getTime() - new Date(value).getTime();

  if (elapsed < 5000) {
    return "刚刚";
  }

  var i = 0;

  while (i + 2 < thresholds.length && elapsed * 1.1 > thresholds[i + 2]) {
    i += 2;
  }

  var divisor = thresholds[i];
  var text = thresholds[i + 1];
  var units = Math.round(elapsed / divisor);

  if (units > 3 && i === thresholds.length - 2) {
    return date.toLocaleDateString(undefined, formatOptions);
  }

  return "" + units + text + "\u524D";
}
},{}],"measure.ts":[function(require,module,exports) {
"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.startMeasuring = startMeasuring;
exports.scheduleMeasure = scheduleMeasure;
var hostOrigin;

function startMeasuring(origin) {
  hostOrigin = origin;
  addEventListener("resize", scheduleMeasure);
  addEventListener("load", scheduleMeasure);
}

var lastHeight = -1;

function measure() {
  var height = document.body.scrollHeight + 60;

  if (height === lastHeight) {
    return;
  }

  lastHeight = height;
  var message = {
    type: "resize",
    height: height
  };
  parent.postMessage(message, hostOrigin);
}

var lastMeasure = 0;

function scheduleMeasure() {
  var now = Date.now();

  if (now - lastMeasure > 50) {
    lastMeasure = now;
    setTimeout(measure, 50);
  }
}
},{}],"timeline-component.ts":[function(require,module,exports) {
"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.KBTimeLineComponent = void 0;

var _commentComponent = require("./comment-component");

var _timeAgo = require("./time-ago");

var _measure = require("./measure");

var KBTimeLineComponent = function () {
  function KBTimeLineComponent(config) {
    var _this = this;

    this.config = config;
    this.page = 1;
    this.loading = false;
    this.element = document.createElement("div");
    this.element.className = "kb-comment-list";
    this.commentComp = new _commentComponent.KBCommentComponent(this.config);
    this.moreDOM = document.createElement("div");
    this.moreDOM.style.textAlign = "center";
    this.moreDOM.innerHTML = "<button class=\"more-btn\">\u52A0\u8F7D\u66F4\u591A</button>";
    this.moreDOM.querySelector(".more-btn").addEventListener("click", function () {
      _this.page++;

      _this.getList();
    });
    this.element.addEventListener("click", function (event) {
      var _a, _b, _c;

      var target = event.target;

      if (target.className == "reply-btn") {
        if ((_a = target.parentNode) === null || _a === void 0 ? void 0 : _a.contains(_this.commentComp.element)) {
          (_b = target.parentNode) === null || _b === void 0 ? void 0 : _b.removeChild(_this.commentComp.element);
        } else {
          (_c = target.parentNode) === null || _c === void 0 ? void 0 : _c.appendChild(_this.commentComp.element);

          _this.commentComp.setReply(target.dataset.pid || "", target.dataset.rid || "");

          _this.commentComp.setEvent(function () {
            var _a;

            (_a = target.parentNode) === null || _a === void 0 ? void 0 : _a.removeChild(_this.commentComp.element);

            _this.getList();
          });
        }

        (0, _measure.scheduleMeasure)();
      }
    });
  }

  KBTimeLineComponent.prototype.getList = function (page) {
    var _this = this;

    if (this.loading == true) return;
    this.loading = true;
    page = page || this.page || 1;

    if (page == 1) {
      this.element.innerHTML = "";
    }

    var loadingDOM = document.createElement("div");
    loadingDOM.innerHTML = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" style=\"margin: auto; display: block; shape-rendering: auto;\" width=\"80px\" height=\"80px\" viewBox=\"0 0 100 100\" preserveAspectRatio=\"xMidYMid\">\n    <circle cx=\"30\" cy=\"50\" fill=\"#1d3f72\" r=\"20\">\n      <animate attributeName=\"cx\" repeatCount=\"indefinite\" dur=\"1s\" keyTimes=\"0;0.5;1\" values=\"30;70;30\" begin=\"-0.5s\"></animate>\n    </circle>\n    <circle cx=\"70\" cy=\"50\" fill=\"#5699d2\" r=\"20\">\n      <animate attributeName=\"cx\" repeatCount=\"indefinite\" dur=\"1s\" keyTimes=\"0;0.5;1\" values=\"30;70;30\" begin=\"0s\"></animate>\n    </circle>\n    <circle cx=\"30\" cy=\"50\" fill=\"#1d3f72\" r=\"20\">\n      <animate attributeName=\"cx\" repeatCount=\"indefinite\" dur=\"1s\" keyTimes=\"0;0.5;1\" values=\"30;70;30\" begin=\"-0.5s\"></animate>\n      <animate attributeName=\"fill-opacity\" values=\"0;0;1;1\" calcMode=\"discrete\" keyTimes=\"0;0.499;0.5;1\" dur=\"1s\" repeatCount=\"indefinite\"></animate>\n    </circle>\n    </svg>";
    this.element.appendChild(loadingDOM);
    var params = {
      token: this.config.token,
      page: page
    };
    axios.get("api/page", {
      params: params
    }).then(function (res) {
      _this.loading = false;

      _this.element.removeChild(loadingDOM);

      if (res.data.code == 200) {
        var data = res.data.data;
        var pageDOM = document.createElement("div");
        pageDOM.innerHTML = _this.renderCommentItem(data.records);

        _this.element.appendChild(pageDOM);

        if (data.page * data.pageSize < data.total) {
          _this.element.appendChild(_this.moreDOM);
        }

        (0, _measure.scheduleMeasure)();
      }
    });
  };

  KBTimeLineComponent.prototype.renderCommentItem = function (list, first) {
    var _this = this;

    if (first === void 0) {
      first = 0;
    }

    return list.reduce(function (html, item) {
      var pid = first == 0 ? item.id : first;
      html += "\n\t\t\t\t<div class=\"comment-item\">\n\t\t\t\t\t<div class=\"comment-avatar\">\n\t\t\t\t\t\t<img src=\"https://s.gravatar.com/avatar/" + md5(item.mail) + "?s=50&d=retro&r=g\" />\n\t\t\t\t\t</div>\n\t\t\t\t\t<div class=\"comment-message clear-right\">\n\t\t\t\t\t\t<div>\n\t\t\t\t\t\t\t<div class=\"comment-time\">" + (0, _timeAgo.timeAgo)(item.createdAt) + "</div>\n\t\t\t\t\t\t\t<div class=\"comment-nickname\"><a target=\"_black\" href=\"" + item.site + "\">" + item.nickName + "</a></div>\n\t\t\t\t\t\t</div>\n            <div class=\"comment-content\">" + item.content + "</div>\n            <div class=\"comment-option\"><a class=\"reply-btn\" data-rid=\"" + item.id + "\" data-pid=\"" + pid + "\" href=\"javascript:\">\u56DE\u590D</a></div>\n\t\t\t\t\t</div>\n\t\t\t\t\t<div class=\"comment-replys\">\n\t\t\t\t\t\t" + (Array.isArray(item.replys) ? _this.renderCommentItem(item.replys, pid) : "") + "\n\t\t\t\t\t</div>\n\t\t\t\t</div>";
      return html;
    }, "");
  };

  return KBTimeLineComponent;
}();

exports.KBTimeLineComponent = KBTimeLineComponent;
},{"./comment-component":"comment-component.ts","./time-ago":"time-ago.ts","./measure":"measure.ts"}],"deparam.ts":[function(require,module,exports) {
"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.deparam = deparam;
exports.param = param;

function deparam(query) {
  var match;
  var plus = /\+/g;
  var search = /([^&=]+)=?([^&]*)/g;

  var decode = function decode(s) {
    return decodeURIComponent(s.replace(plus, ' '));
  };

  var params = {};

  while (match = search.exec(query)) {
    params[decode(match[1])] = decode(match[2]);
  }

  return params;
}

function param(obj) {
  var parts = [];

  for (var name in obj) {
    if (obj.hasOwnProperty(name) && obj[name]) {
      parts.push(encodeURIComponent(name) + "=" + encodeURIComponent(obj[name]));
    }
  }

  return parts.join('&');
}
},{}],"comment.ts":[function(require,module,exports) {
"use strict";

var _commentComponent = require("./comment-component");

var _timelineComponent = require("./timeline-component");

var _deparam = require("./deparam");

var _measure = require("./measure");

var KBComment = function () {
  function KBComment() {
    var _this = this;

    this.config = {
      theme: "light",
      token: ""
    };
    this.container = document.createElement("div");
    document.body.appendChild(this.container);
    var params = (0, _deparam.deparam)(location.search.replace("?", ""));

    if (!this.container) {
      console.error("未设定渲染容器");
    }

    (0, _measure.startMeasuring)(params.origin);
    this.container.className = "kb-comment-container";
    this.config.theme = params.theme || "light";
    this.config.token = params.token || location.pathname;
    this.load().then(function () {
      var _a, _b;

      _this.commentComponent = new _commentComponent.KBCommentComponent(_this.config);
      (_a = _this.container) === null || _a === void 0 ? void 0 : _a.appendChild(_this.commentComponent.element);
      _this.timelineComponent = new _timelineComponent.KBTimeLineComponent(_this.config);
      (_b = _this.container) === null || _b === void 0 ? void 0 : _b.appendChild(_this.timelineComponent.element);

      _this.timelineComponent.getList();

      var self = _this;

      _this.commentComponent.setEvent(function () {
        var _a;

        (_a = self.timelineComponent) === null || _a === void 0 ? void 0 : _a.getList();
      });
    });
  }

  KBComment.prototype.load = function () {
    return Promise.all([this.loadTheme()]);
  };

  KBComment.prototype.loadTheme = function () {
    var _this = this;

    return new Promise(function (resolve) {
      var link = document.createElement("link");
      link.rel = "stylesheet";
      link.setAttribute("crossorigin", "anonymous");
      link.onload = resolve;
      link.href = "./static/themes/" + _this.config.theme + ".css";
      document.head.appendChild(link);
    });
  };

  return KBComment;
}();

new KBComment();
},{"./comment-component":"comment-component.ts","./timeline-component":"timeline-component.ts","./deparam":"deparam.ts","./measure":"measure.ts"}],"C:/Users/guoren/AppData/Roaming/npm/node_modules/parcel-bundler/src/builtins/hmr-runtime.js":[function(require,module,exports) {
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
  var ws = new WebSocket(protocol + '://' + hostname + ':' + "64735" + '/');

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
},{}]},{},["C:/Users/guoren/AppData/Roaming/npm/node_modules/parcel-bundler/src/builtins/hmr-runtime.js","comment.ts"], null)
//# sourceMappingURL=/comment.25536dae.js.map