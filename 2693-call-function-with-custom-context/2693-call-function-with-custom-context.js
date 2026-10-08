Function.prototype.callPolyfill = function(context, ...args) {
    const sym = Symbol();
    context[sym] = this;
    const result = context[sym](...args);
    delete context[sym];
    return result;
};