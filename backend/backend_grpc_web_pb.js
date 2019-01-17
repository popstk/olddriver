/**
 * @fileoverview gRPC-Web generated client stub for backend
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.backend = require('./backend_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.backend.SpiderClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.backend.SpiderPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.backend.SpiderClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.backend.SpiderClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.backend.SearchRequest,
 *   !proto.backend.SearchReply>}
 */
const methodInfo_Spider_Search = new grpc.web.AbstractClientBase.MethodInfo(
  proto.backend.SearchReply,
  /** @param {!proto.backend.SearchRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.backend.SearchReply.deserializeBinary
);


/**
 * @param {!proto.backend.SearchRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.backend.SearchReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.backend.SearchReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.backend.SpiderClient.prototype.search =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/backend.Spider/Search',
      request,
      metadata,
      methodInfo_Spider_Search,
      callback);
};


/**
 * @param {!proto.backend.SearchRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.backend.SearchReply>}
 *     The XHR Node Readable Stream
 */
proto.backend.SpiderPromiseClient.prototype.search =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.search(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.backend;

