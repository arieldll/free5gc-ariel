/**
 * The IndividualSubscriptionDocumentController file is a very simple one, which does not need to be changed manually,
 * unless there's a case where business logic routes the request to an entity which is not
 * the service.
 * The heavy lifting of the Controller item is done in Request.js - that is where request
 * parameters are extracted and sent to the service, and where response is handled.
 */

const Controller = require('./Controller');
const service = require('../services/IndividualSubscriptionDocumentService');
const deleteIndividualSubcription = async (request, response) => {
  await Controller.handleRequest(request, response, service.deleteIndividualSubcription);
};

const getIndividualSubcription = async (request, response) => {
  await Controller.handleRequest(request, response, service.getIndividualSubcription);
};

const replaceIndividualSubcription = async (request, response) => {
  await Controller.handleRequest(request, response, service.replaceIndividualSubcription);
};


module.exports = {
  deleteIndividualSubcription,
  getIndividualSubcription,
  replaceIndividualSubcription,
};
