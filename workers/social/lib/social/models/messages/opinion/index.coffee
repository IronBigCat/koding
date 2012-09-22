JPost = require '../post'

module.exports = class JOpinion extends JPost

  # @mixin Followable
  # @::mixin Followable::
  # @::mixin Taggable::
  # @::mixin Notifying::
  # @mixin Flaggable
  # @::mixin Flaggable::
  # @::mixin Likeable::

  {Base,ObjectId,ObjectRef,secure,dash,daisy} = require 'bongo'
  {Relationship, Message} = require 'jraphical'
  {extend} = require 'underscore'

  {log} = console

  @share()

  schema = extend {}, Message.schema, {
    isLowQuality  : Boolean
    counts        :
      followers   :
        type      : Number
        default   : 0
      following   :
        type      : Number
        default   : 0
    originType  :
      type      : String
      required  : yes
    originId    :
      type      : ObjectId
      required  : yes
    deletedAt   : Date
    deletedBy   : ObjectRef
    meta        : require 'bongo/bundles/meta'
  }

  @set
    emitFollowingActivities: yes
    taggedContentRole : 'reply'
    tagRole           : 'tag'
    sharedMethods : JPost.sharedMethods
    schema        : schema
    relationships     :
      comment         :
        type          : "JComment"
        as            : 'reply'
      participant     :
        targetType    : "JAccount"
        as            : ['author','commenter']
      likedBy         :
        targetType    : "JAccount"
        as            : 'like'
      repliesActivity :
        targetType    : "CRepliesActivity"
        as            : 'repliesActivity'
      tag             :
        targetType    : "JTag"
        as            : 'tag'
      follower        :
        as            : 'follower'
        targetType    : "JAccount"

  @getActivityType =-> require './opinionactivity'

  @getAuthorType =-> require '../../account'

  @getFlagRole =-> ['sender', 'recipient']

  createKodingError =(err)->
    kodingErr = new KodingError(err.message)
    for own prop of err
      kodingErr[prop] = err[prop]
    kodingErr

  @create = secure (client, data, callback)->
    codeSnip =
      title       : data.title
      body        : data.body
      meta        : data.meta
    JPost.create.call @, client, codeSnip, callback

  delete: secure ({connection:{delegate}}, callback)->
    originId = @getAt 'originId'
    unless delegate.getId().equals originId
      callback new KodingError 'Access denied!'
    else
      id = @getId()
      {getDeleteHelper} = Relationship
      rel = null
      message = null

      queue = [
        ->
          Relationship.one {
            targetId    : id
            as          : "opinion"
          }, (err, rel_)->
            if err
              callback err
            else
              rel = rel_
              queue.next(err)
        ->
          rel.fetchSource (err, message_)->
            if err
              callback err
            else
              message = message_
              queue.next(err)
        ->
          message.removeReply rel, (err)-> queue.next(err)

        getDeleteHelper {
          targetId    : id
          sourceName  : /Activity$/
        }, 'source', (err)-> queue.next(err)

        getDeleteHelper {
          targetName  : {$ne : 'JAccount'}
          sourceId    : id
          sourceName  : 'JOpinion'
        }, 'target', (err)-> queue.next(err)

        ->
          Relationship.remove {
            targetId  : id
            as        : 'opinion'
          }, (err)-> queue.next(err)
        =>
          @remove -> queue.next()
        =>
          @emit "OpinionIsDeleted", 1
          callback null
      ]
      daisy queue

  modify: secure (client, data, callback)->
    opinion =
      title       : data.title
      body        : data.body
      meta        : data.meta
    JPost::modify.call @, client, opinion, callback

  reply: secure (client, comment, callback)->
    JPost::reply.call @, client, JComment, comment, callback