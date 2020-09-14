package gen

type key int

// consts ...
const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  task(id: ID, q: String, filter: TaskFilterType): Task
  tasks(offset: Int, limit: Int = 30, q: String, sort: [TaskSortType!], filter: TaskFilterType): TaskResultType!
  taskCategory(id: ID, q: String, filter: TaskCategoryFilterType): TaskCategory
  taskCategories(offset: Int, limit: Int = 30, q: String, sort: [TaskCategorySortType!], filter: TaskCategoryFilterType): TaskCategoryResultType!
  company(id: ID, q: String, filter: CompanyFilterType): Company
  companies(offset: Int, limit: Int = 30, q: String, sort: [CompanySortType!], filter: CompanyFilterType): CompanyResultType!
  user(id: ID, q: String, filter: UserFilterType): User
  users(offset: Int, limit: Int = 30, q: String, sort: [UserSortType!], filter: UserFilterType): UserResultType!
  plainEntity(id: ID, q: String, filter: PlainEntityFilterType): PlainEntity
  plainEntities(offset: Int, limit: Int = 30, q: String, sort: [PlainEntitySortType!], filter: PlainEntityFilterType): PlainEntityResultType!
}

type Mutation {
  createTask(input: TaskCreateInput!): Task!
  updateTask(id: ID!, input: TaskUpdateInput!): Task!
  deleteTask(id: ID!): Task!
  deleteAllTasks: Boolean!
  createTaskCategory(input: TaskCategoryCreateInput!): TaskCategory!
  updateTaskCategory(id: ID!, input: TaskCategoryUpdateInput!): TaskCategory!
  deleteTaskCategory(id: ID!): TaskCategory!
  deleteAllTaskCategories: Boolean!
  createCompany(input: CompanyCreateInput!): Company!
  updateCompany(id: ID!, input: CompanyUpdateInput!): Company!
  deleteCompany(id: ID!): Company!
  deleteAllCompanies: Boolean!
  createUser(input: UserCreateInput!): User!
  updateUser(id: ID!, input: UserUpdateInput!): User!
  deleteUser(id: ID!): User!
  deleteAllUsers: Boolean!
  createPlainEntity(input: PlainEntityCreateInput!): PlainEntity!
  updatePlainEntity(id: ID!, input: PlainEntityUpdateInput!): PlainEntity!
  deletePlainEntity(id: ID!): PlainEntity!
  deleteAllPlainEntities: Boolean!
}

enum ObjectSortType {
  ASC
  DESC
}

enum TaskState {
  CREATED
  IN_PROGRESS
  RESOLVED
}

type TaskMeta {
  key: String!
  value: String
}

type Task {
  id: ID!
  title: String
  completed: Boolean
  state: TaskState
  dueDate: Time
  metas: [TaskMeta!]
  meta: TaskMeta
  assignee: User
  owner: User!
  parentTask: Task
  subtasks: [Task!]!
  categories: [TaskCategory!]!
  assigneeId: ID
  ownerId: ID
  parentTaskId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  subtasksIds: [ID!]!
  subtasksConnection(offset: Int, limit: Int = 30, q: String, sort: [TaskSortType!], filter: TaskFilterType): TaskResultType!
  categoriesIds: [ID!]!
  categoriesConnection(offset: Int, limit: Int = 30, q: String, sort: [TaskCategorySortType!], filter: TaskCategoryFilterType): TaskCategoryResultType!
}

type TaskCategory {
  id: ID!
  name: String
  tasks: [Task!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  tasksIds: [ID!]!
  tasksConnection(offset: Int, limit: Int = 30, q: String, sort: [TaskSortType!], filter: TaskFilterType): TaskResultType!
}

extend type Query {
  hello: String!
  topCompanies: [Company!]!
}

interface NamedEntity {
  name: String
}

type Company implements NamedEntity @key(fields: "id") {
  id: ID!
  name: String
  countryId: ID
  country: Country
  employees: [User!]!
  reviews: [Review!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  employeesIds: [ID!]!
  employeesConnection(offset: Int, limit: Int = 30, q: String, sort: [UserSortType!], filter: UserFilterType): UserResultType!
}

extend type Company {
  uppercaseName: String!
}

type Address {
  street: String
  city: String
  zip: String
}

type User {
  id: ID!
  code: Int
  email: String
  firstName: String
  lastName: String
  addressRaw: String
  address: Address
  salary: Int
  employers: [Company!]!
  tasks: [Task!]!
  createdTasks: [Task!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  employersIds: [ID!]!
  employersConnection(offset: Int, limit: Int = 30, q: String, sort: [CompanySortType!], filter: CompanyFilterType): CompanyResultType!
  tasksIds: [ID!]!
  tasksConnection(offset: Int, limit: Int = 30, q: String, sort: [TaskSortType!], filter: TaskFilterType): TaskResultType!
  createdTasksIds: [ID!]!
  createdTasksConnection(offset: Int, limit: Int = 30, q: String, sort: [TaskSortType!], filter: TaskFilterType): TaskResultType!
}

extend type Review @entity @key(fields: "id") {
  id: ID! @external
  referenceID: ID! @external
  company: Company @requires(fields: "referenceID")
}

extend type Country @entity @key(fields: "id") {
  id: ID! @external
}

type PlainEntity {
  id: ID!
  date: Time
  text: String
  shortText: String!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

input TaskCreateInput {
  id: ID
  title: String
  completed: Boolean
  state: TaskState
  dueDate: Time
  metas: [TaskMetaInput]
  meta: TaskMetaInput
  assigneeId: ID
  ownerId: ID
  parentTaskId: ID
  subtasksIds: [ID!]
  categoriesIds: [ID!]
}

input TaskUpdateInput {
  title: String
  completed: Boolean
  state: TaskState
  dueDate: Time
  metas: [TaskMetaInput]
  meta: TaskMetaInput
  assigneeId: ID
  ownerId: ID
  parentTaskId: ID
  subtasksIds: [ID!]
  categoriesIds: [ID!]
}

input TaskSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  title: ObjectSortType
  titleMin: ObjectSortType
  titleMax: ObjectSortType
  completed: ObjectSortType
  completedMin: ObjectSortType
  completedMax: ObjectSortType
  state: ObjectSortType
  stateMin: ObjectSortType
  stateMax: ObjectSortType
  dueDate: ObjectSortType
  dueDateMin: ObjectSortType
  dueDateMax: ObjectSortType
  metas: ObjectSortType
  metasMin: ObjectSortType
  metasMax: ObjectSortType
  meta: ObjectSortType
  metaMin: ObjectSortType
  metaMax: ObjectSortType
  assigneeId: ObjectSortType
  assigneeIdMin: ObjectSortType
  assigneeIdMax: ObjectSortType
  ownerId: ObjectSortType
  ownerIdMin: ObjectSortType
  ownerIdMax: ObjectSortType
  parentTaskId: ObjectSortType
  parentTaskIdMin: ObjectSortType
  parentTaskIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  subtasksIds: ObjectSortType
  subtasksIdsMin: ObjectSortType
  subtasksIdsMax: ObjectSortType
  categoriesIds: ObjectSortType
  categoriesIdsMin: ObjectSortType
  categoriesIdsMax: ObjectSortType
  assignee: UserSortType
  owner: UserSortType
  parentTask: TaskSortType
  subtasks: TaskSortType
  categories: TaskCategorySortType
}

input TaskFilterType {
  AND: [TaskFilterType!]
  OR: [TaskFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  title: String
  titleMin: String
  titleMax: String
  title_ne: String
  titleMin_ne: String
  titleMax_ne: String
  title_gt: String
  titleMin_gt: String
  titleMax_gt: String
  title_lt: String
  titleMin_lt: String
  titleMax_lt: String
  title_gte: String
  titleMin_gte: String
  titleMax_gte: String
  title_lte: String
  titleMin_lte: String
  titleMax_lte: String
  title_in: [String!]
  titleMin_in: [String!]
  titleMax_in: [String!]
  title_like: String
  titleMin_like: String
  titleMax_like: String
  title_prefix: String
  titleMin_prefix: String
  titleMax_prefix: String
  title_suffix: String
  titleMin_suffix: String
  titleMax_suffix: String
  title_null: Boolean
  completed: Boolean
  completedMin: Boolean
  completedMax: Boolean
  completed_ne: Boolean
  completedMin_ne: Boolean
  completedMax_ne: Boolean
  completed_gt: Boolean
  completedMin_gt: Boolean
  completedMax_gt: Boolean
  completed_lt: Boolean
  completedMin_lt: Boolean
  completedMax_lt: Boolean
  completed_gte: Boolean
  completedMin_gte: Boolean
  completedMax_gte: Boolean
  completed_lte: Boolean
  completedMin_lte: Boolean
  completedMax_lte: Boolean
  completed_in: [Boolean!]
  completedMin_in: [Boolean!]
  completedMax_in: [Boolean!]
  completed_null: Boolean
  state: TaskState
  stateMin: TaskState
  stateMax: TaskState
  state_ne: TaskState
  stateMin_ne: TaskState
  stateMax_ne: TaskState
  state_gt: TaskState
  stateMin_gt: TaskState
  stateMax_gt: TaskState
  state_lt: TaskState
  stateMin_lt: TaskState
  stateMax_lt: TaskState
  state_gte: TaskState
  stateMin_gte: TaskState
  stateMax_gte: TaskState
  state_lte: TaskState
  stateMin_lte: TaskState
  stateMax_lte: TaskState
  state_in: [TaskState!]
  stateMin_in: [TaskState!]
  stateMax_in: [TaskState!]
  state_null: Boolean
  dueDate: Time
  dueDateMin: Time
  dueDateMax: Time
  dueDate_ne: Time
  dueDateMin_ne: Time
  dueDateMax_ne: Time
  dueDate_gt: Time
  dueDateMin_gt: Time
  dueDateMax_gt: Time
  dueDate_lt: Time
  dueDateMin_lt: Time
  dueDateMax_lt: Time
  dueDate_gte: Time
  dueDateMin_gte: Time
  dueDateMax_gte: Time
  dueDate_lte: Time
  dueDateMin_lte: Time
  dueDateMax_lte: Time
  dueDate_in: [Time!]
  dueDateMin_in: [Time!]
  dueDateMax_in: [Time!]
  dueDate_null: Boolean
  assigneeId: ID
  assigneeIdMin: ID
  assigneeIdMax: ID
  assigneeId_ne: ID
  assigneeIdMin_ne: ID
  assigneeIdMax_ne: ID
  assigneeId_gt: ID
  assigneeIdMin_gt: ID
  assigneeIdMax_gt: ID
  assigneeId_lt: ID
  assigneeIdMin_lt: ID
  assigneeIdMax_lt: ID
  assigneeId_gte: ID
  assigneeIdMin_gte: ID
  assigneeIdMax_gte: ID
  assigneeId_lte: ID
  assigneeIdMin_lte: ID
  assigneeIdMax_lte: ID
  assigneeId_in: [ID!]
  assigneeIdMin_in: [ID!]
  assigneeIdMax_in: [ID!]
  assigneeId_null: Boolean
  ownerId: ID
  ownerIdMin: ID
  ownerIdMax: ID
  ownerId_ne: ID
  ownerIdMin_ne: ID
  ownerIdMax_ne: ID
  ownerId_gt: ID
  ownerIdMin_gt: ID
  ownerIdMax_gt: ID
  ownerId_lt: ID
  ownerIdMin_lt: ID
  ownerIdMax_lt: ID
  ownerId_gte: ID
  ownerIdMin_gte: ID
  ownerIdMax_gte: ID
  ownerId_lte: ID
  ownerIdMin_lte: ID
  ownerIdMax_lte: ID
  ownerId_in: [ID!]
  ownerIdMin_in: [ID!]
  ownerIdMax_in: [ID!]
  ownerId_null: Boolean
  parentTaskId: ID
  parentTaskIdMin: ID
  parentTaskIdMax: ID
  parentTaskId_ne: ID
  parentTaskIdMin_ne: ID
  parentTaskIdMax_ne: ID
  parentTaskId_gt: ID
  parentTaskIdMin_gt: ID
  parentTaskIdMax_gt: ID
  parentTaskId_lt: ID
  parentTaskIdMin_lt: ID
  parentTaskIdMax_lt: ID
  parentTaskId_gte: ID
  parentTaskIdMin_gte: ID
  parentTaskIdMax_gte: ID
  parentTaskId_lte: ID
  parentTaskIdMin_lte: ID
  parentTaskIdMax_lte: ID
  parentTaskId_in: [ID!]
  parentTaskIdMin_in: [ID!]
  parentTaskIdMax_in: [ID!]
  parentTaskId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
  assignee: UserFilterType
  owner: UserFilterType
  parentTask: TaskFilterType
  subtasks: TaskFilterType
  categories: TaskCategoryFilterType
}

type TaskResultType {
  items: [Task!]!
  count: Int!
}

input TaskCategoryCreateInput {
  id: ID
  name: String
  tasksIds: [ID!]
}

input TaskCategoryUpdateInput {
  name: String
  tasksIds: [ID!]
}

input TaskCategorySortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  tasksIds: ObjectSortType
  tasksIdsMin: ObjectSortType
  tasksIdsMax: ObjectSortType
  tasks: TaskSortType
}

input TaskCategoryFilterType {
  AND: [TaskCategoryFilterType!]
  OR: [TaskCategoryFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
  tasks: TaskFilterType
}

type TaskCategoryResultType {
  items: [TaskCategory!]!
  count: Int!
}

input CompanyCreateInput {
  id: ID
  name: String
  countryId: ID
  employeesIds: [ID!]
}

input CompanyUpdateInput {
  name: String
  countryId: ID
  employeesIds: [ID!]
}

input CompanySortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  countryId: ObjectSortType
  countryIdMin: ObjectSortType
  countryIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  employeesIds: ObjectSortType
  employeesIdsMin: ObjectSortType
  employeesIdsMax: ObjectSortType
  employees: UserSortType
}

input CompanyFilterType {
  AND: [CompanyFilterType!]
  OR: [CompanyFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  countryId: ID
  countryIdMin: ID
  countryIdMax: ID
  countryId_ne: ID
  countryIdMin_ne: ID
  countryIdMax_ne: ID
  countryId_gt: ID
  countryIdMin_gt: ID
  countryIdMax_gt: ID
  countryId_lt: ID
  countryIdMin_lt: ID
  countryIdMax_lt: ID
  countryId_gte: ID
  countryIdMin_gte: ID
  countryIdMax_gte: ID
  countryId_lte: ID
  countryIdMin_lte: ID
  countryIdMax_lte: ID
  countryId_in: [ID!]
  countryIdMin_in: [ID!]
  countryIdMax_in: [ID!]
  countryId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
  employees: UserFilterType
}

type CompanyResultType {
  items: [Company!]!
  count: Int!
}

input UserCreateInput {
  id: ID
  code: Int
  email: String
  firstName: String
  lastName: String
  addressRaw: String
  salary: Int
  employersIds: [ID!]
  tasksIds: [ID!]
  createdTasksIds: [ID!]
}

input UserUpdateInput {
  code: Int
  email: String
  firstName: String
  lastName: String
  addressRaw: String
  salary: Int
  employersIds: [ID!]
  tasksIds: [ID!]
  createdTasksIds: [ID!]
}

input UserSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  code: ObjectSortType
  codeMin: ObjectSortType
  codeMax: ObjectSortType
  codeAvg: ObjectSortType
  email: ObjectSortType
  emailMin: ObjectSortType
  emailMax: ObjectSortType
  firstName: ObjectSortType
  firstNameMin: ObjectSortType
  firstNameMax: ObjectSortType
  lastName: ObjectSortType
  lastNameMin: ObjectSortType
  lastNameMax: ObjectSortType
  addressRaw: ObjectSortType
  addressRawMin: ObjectSortType
  addressRawMax: ObjectSortType
  salary: ObjectSortType
  salaryMin: ObjectSortType
  salaryMax: ObjectSortType
  salaryAvg: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  employersIds: ObjectSortType
  employersIdsMin: ObjectSortType
  employersIdsMax: ObjectSortType
  tasksIds: ObjectSortType
  tasksIdsMin: ObjectSortType
  tasksIdsMax: ObjectSortType
  createdTasksIds: ObjectSortType
  createdTasksIdsMin: ObjectSortType
  createdTasksIdsMax: ObjectSortType
  employers: CompanySortType
  tasks: TaskSortType
  createdTasks: TaskSortType
}

input UserFilterType {
  AND: [UserFilterType!]
  OR: [UserFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  code: Int
  codeMin: Int
  codeMax: Int
  codeAvg: Int
  code_ne: Int
  codeMin_ne: Int
  codeMax_ne: Int
  codeAvg_ne: Int
  code_gt: Int
  codeMin_gt: Int
  codeMax_gt: Int
  codeAvg_gt: Int
  code_lt: Int
  codeMin_lt: Int
  codeMax_lt: Int
  codeAvg_lt: Int
  code_gte: Int
  codeMin_gte: Int
  codeMax_gte: Int
  codeAvg_gte: Int
  code_lte: Int
  codeMin_lte: Int
  codeMax_lte: Int
  codeAvg_lte: Int
  code_in: [Int!]
  codeMin_in: [Int!]
  codeMax_in: [Int!]
  codeAvg_in: [Int!]
  code_null: Boolean
  email: String
  emailMin: String
  emailMax: String
  email_ne: String
  emailMin_ne: String
  emailMax_ne: String
  email_gt: String
  emailMin_gt: String
  emailMax_gt: String
  email_lt: String
  emailMin_lt: String
  emailMax_lt: String
  email_gte: String
  emailMin_gte: String
  emailMax_gte: String
  email_lte: String
  emailMin_lte: String
  emailMax_lte: String
  email_in: [String!]
  emailMin_in: [String!]
  emailMax_in: [String!]
  email_like: String
  emailMin_like: String
  emailMax_like: String
  email_prefix: String
  emailMin_prefix: String
  emailMax_prefix: String
  email_suffix: String
  emailMin_suffix: String
  emailMax_suffix: String
  email_null: Boolean
  firstName: String
  firstNameMin: String
  firstNameMax: String
  firstName_ne: String
  firstNameMin_ne: String
  firstNameMax_ne: String
  firstName_gt: String
  firstNameMin_gt: String
  firstNameMax_gt: String
  firstName_lt: String
  firstNameMin_lt: String
  firstNameMax_lt: String
  firstName_gte: String
  firstNameMin_gte: String
  firstNameMax_gte: String
  firstName_lte: String
  firstNameMin_lte: String
  firstNameMax_lte: String
  firstName_in: [String!]
  firstNameMin_in: [String!]
  firstNameMax_in: [String!]
  firstName_like: String
  firstNameMin_like: String
  firstNameMax_like: String
  firstName_prefix: String
  firstNameMin_prefix: String
  firstNameMax_prefix: String
  firstName_suffix: String
  firstNameMin_suffix: String
  firstNameMax_suffix: String
  firstName_null: Boolean
  lastName: String
  lastNameMin: String
  lastNameMax: String
  lastName_ne: String
  lastNameMin_ne: String
  lastNameMax_ne: String
  lastName_gt: String
  lastNameMin_gt: String
  lastNameMax_gt: String
  lastName_lt: String
  lastNameMin_lt: String
  lastNameMax_lt: String
  lastName_gte: String
  lastNameMin_gte: String
  lastNameMax_gte: String
  lastName_lte: String
  lastNameMin_lte: String
  lastNameMax_lte: String
  lastName_in: [String!]
  lastNameMin_in: [String!]
  lastNameMax_in: [String!]
  lastName_like: String
  lastNameMin_like: String
  lastNameMax_like: String
  lastName_prefix: String
  lastNameMin_prefix: String
  lastNameMax_prefix: String
  lastName_suffix: String
  lastNameMin_suffix: String
  lastNameMax_suffix: String
  lastName_null: Boolean
  addressRaw: String
  addressRawMin: String
  addressRawMax: String
  addressRaw_ne: String
  addressRawMin_ne: String
  addressRawMax_ne: String
  addressRaw_gt: String
  addressRawMin_gt: String
  addressRawMax_gt: String
  addressRaw_lt: String
  addressRawMin_lt: String
  addressRawMax_lt: String
  addressRaw_gte: String
  addressRawMin_gte: String
  addressRawMax_gte: String
  addressRaw_lte: String
  addressRawMin_lte: String
  addressRawMax_lte: String
  addressRaw_in: [String!]
  addressRawMin_in: [String!]
  addressRawMax_in: [String!]
  addressRaw_like: String
  addressRawMin_like: String
  addressRawMax_like: String
  addressRaw_prefix: String
  addressRawMin_prefix: String
  addressRawMax_prefix: String
  addressRaw_suffix: String
  addressRawMin_suffix: String
  addressRawMax_suffix: String
  addressRaw_null: Boolean
  salary: Int
  salaryMin: Int
  salaryMax: Int
  salaryAvg: Int
  salary_ne: Int
  salaryMin_ne: Int
  salaryMax_ne: Int
  salaryAvg_ne: Int
  salary_gt: Int
  salaryMin_gt: Int
  salaryMax_gt: Int
  salaryAvg_gt: Int
  salary_lt: Int
  salaryMin_lt: Int
  salaryMax_lt: Int
  salaryAvg_lt: Int
  salary_gte: Int
  salaryMin_gte: Int
  salaryMax_gte: Int
  salaryAvg_gte: Int
  salary_lte: Int
  salaryMin_lte: Int
  salaryMax_lte: Int
  salaryAvg_lte: Int
  salary_in: [Int!]
  salaryMin_in: [Int!]
  salaryMax_in: [Int!]
  salaryAvg_in: [Int!]
  salary_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
  employers: CompanyFilterType
  tasks: TaskFilterType
  createdTasks: TaskFilterType
}

type UserResultType {
  items: [User!]!
  count: Int!
}

input PlainEntityCreateInput {
  id: ID
  date: Time
  text: String
}

input PlainEntityUpdateInput {
  date: Time
  text: String
}

input PlainEntitySortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  date: ObjectSortType
  dateMin: ObjectSortType
  dateMax: ObjectSortType
  text: ObjectSortType
  textMin: ObjectSortType
  textMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
}

input PlainEntityFilterType {
  AND: [PlainEntityFilterType!]
  OR: [PlainEntityFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  date: Time
  dateMin: Time
  dateMax: Time
  date_ne: Time
  dateMin_ne: Time
  dateMax_ne: Time
  date_gt: Time
  dateMin_gt: Time
  dateMax_gt: Time
  date_lt: Time
  dateMin_lt: Time
  dateMax_lt: Time
  date_gte: Time
  dateMin_gte: Time
  dateMax_gte: Time
  date_lte: Time
  dateMin_lte: Time
  dateMax_lte: Time
  date_in: [Time!]
  dateMin_in: [Time!]
  dateMax_in: [Time!]
  date_null: Boolean
  text: String
  textMin: String
  textMax: String
  text_ne: String
  textMin_ne: String
  textMax_ne: String
  text_gt: String
  textMin_gt: String
  textMax_gt: String
  text_lt: String
  textMin_lt: String
  textMax_lt: String
  text_gte: String
  textMin_gte: String
  textMax_gte: String
  text_lte: String
  textMin_lte: String
  textMax_lte: String
  text_in: [String!]
  textMin_in: [String!]
  textMax_in: [String!]
  text_like: String
  textMin_like: String
  textMax_like: String
  text_prefix: String
  textMin_prefix: String
  textMax_prefix: String
  text_suffix: String
  textMin_suffix: String
  textMax_suffix: String
  text_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
}

type PlainEntityResultType {
  items: [PlainEntity!]!
  count: Int!
}

input TaskMetaInput {
  key: String!
  value: String
}`
)
