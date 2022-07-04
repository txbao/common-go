/*
 Navicat Premium Data Transfer

 Source Server         : Dev-开发环境
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : 10.10.10.159:3306
 Source Schema         : bank-activity

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 29/09/2021 18:22:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for base_activity
-- ----------------------------
DROP TABLE IF EXISTS `base_activity`;
CREATE TABLE `base_activity`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `customer_id` int(11) NOT NULL DEFAULT 0 COMMENT ' 客户ID',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动名称',
  `start_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间',
  `end_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '结束时间',
  `remarks` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '活动描述',
  `admin_id` int(11) NOT NULL DEFAULT 0 COMMENT ' 创建人ID',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '活动类型：1营销活动，2演示活动',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态：1有效、0无效',
  `qr_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '二维码',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `customer_id`(`customer_id`) USING BTREE,
  INDEX `admin_id`(`admin_id`) USING BTREE,
  INDEX `is_del_idx`(`is_del`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '活动表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_activity_config
-- ----------------------------
DROP TABLE IF EXISTS `base_activity_config`;
CREATE TABLE `base_activity_config`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `activity_id` int(11) NOT NULL DEFAULT 0 COMMENT '活动ID',
  `is_banner` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示banner：1是、0否',
  `is_icon` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示icon：1是、0否',
  `is_search_bar` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示搜索栏：1是、0否',
  `is_tab_bar` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示Tab导航栏：1是、0否',
  `is_nav_bar_bottom` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示底部导航栏：1是、0否',
  `is_promote_bar_bottom` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示底部推广栏：1是、0否',
  `share_title` varchar(100) NOT NULL DEFAULT '' COMMENT '分享标题',
  `share_desc` varchar(100) NOT NULL DEFAULT '' COMMENT '分享描述',
  `share_img_url` varchar(255) NOT NULL DEFAULT '' COMMENT '分享图标',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_activity_id`(`activity_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '活动配置表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_banner
-- ----------------------------
DROP TABLE IF EXISTS `base_banner`;
CREATE TABLE `base_banner`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `picture` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `jump_mode` tinyint(1) NOT NULL DEFAULT 0 COMMENT '跳转方式：1外部链接、2自定义页面',
  `url` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转方法=外部链接时，外部跳转地址',
  `topic_id` int(11) NOT NULL DEFAULT 0 COMMENT '跳转方式=自定义页面时，主题ID',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序，降序',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'banner-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_customer
-- ----------------------------
DROP TABLE IF EXISTS `base_customer`;
CREATE TABLE `base_customer`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `salesman_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务员ID',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '客户名称',
  `username` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `password` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '密码',
  `remarks` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `admin_id` int(11) NOT NULL DEFAULT 0 COMMENT ' 创建人ID',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1有效、0无效',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_department
-- ----------------------------
DROP TABLE IF EXISTS `base_department`;
CREATE TABLE `base_department`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1有效、0无效',
  `admin_id` int(11) NOT NULL DEFAULT 0 COMMENT ' 创建人ID',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_icon
-- ----------------------------
DROP TABLE IF EXISTS `base_icon`;
CREATE TABLE `base_icon`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `icon` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'icon',
  `url` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '外部跳转地址',
  `sort` int(11) NULL DEFAULT 0 COMMENT '排序，降序',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'icon-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_module
-- ----------------------------
DROP TABLE IF EXISTS `base_module`;
CREATE TABLE `base_module`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `icon` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'icon',
  `sort` int(11) NULL DEFAULT 0 COMMENT '排序，降序',
  `is_more` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示更多：1开启、0关闭',
  `more_title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更多标题',
  `more_jump_mode` tinyint(1) NOT NULL DEFAULT 0 COMMENT '跳转方式：1外部链接、2自定义页面',
  `more_url` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转方法=外部链接时，URL地址',
  `topic_id` int(11) NOT NULL DEFAULT 0 COMMENT '跳转方式=自定义页面时，主题ID',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '模块-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_module_sub
-- ----------------------------
DROP TABLE IF EXISTS `base_module_sub`;
CREATE TABLE `base_module_sub`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `module_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模块id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块名称',
  `icon` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'icon',
  `sort` int(11) NULL DEFAULT 0 COMMENT '排序，降序',
  `jump_mode` tinyint(1) NOT NULL DEFAULT 0 COMMENT '跳转方式：1外部链接、2自定义页面',
  `url` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转方法=外部链接时，URL地址',
  `topic_id` int(11) NOT NULL DEFAULT 0 COMMENT '跳转方式=自定义页面时，主题ID',
  `style` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '子模块样式 1-单模块 2-双模块',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题-双模块时显示',
  `intro` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '简介-双模块时显示',
  `remarks` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_module_id`(`module_id`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '子模块-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_promote
-- ----------------------------
DROP TABLE IF EXISTS `base_promote`;
CREATE TABLE `base_promote`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `button` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '按钮样式图片',
  `jump_mode` tinyint(1) NOT NULL DEFAULT 0 COMMENT '跳转方式：1外部链接、2自定义页面',
  `url` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转方法=外部链接时，URL地址',
  `topic_id` int(11) NOT NULL DEFAULT 0 COMMENT '跳转方式=自定义页面时，主题ID',
  `picture` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推广位表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_salesman
-- ----------------------------
DROP TABLE IF EXISTS `base_salesman`;
CREATE TABLE `base_salesman`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `dept_id` int(11) NOT NULL DEFAULT 0 COMMENT '部门ID',
  `fullname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `gender` tinyint(1) NOT NULL DEFAULT 1 COMMENT '性别：1男、2女',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1有效、0无效',
  `admin_id` int(11) NOT NULL DEFAULT 0 COMMENT ' 创建人ID',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '业务员表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_topic
-- ----------------------------
DROP TABLE IF EXISTS `base_topic`;
CREATE TABLE `base_topic`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '图文详情',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '主题表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goods_category
-- ----------------------------
DROP TABLE IF EXISTS `goods_category`;
CREATE TABLE `goods_category`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `parent_id` int(11) NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_parent_id`(`parent_id`) USING BTREE,
  INDEX `idx_createdat`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品类目表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goods_sku
-- ----------------------------
DROP TABLE IF EXISTS `goods_sku`;
CREATE TABLE `goods_sku`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `spu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
  `third_sku_id` int(11) NOT NULL DEFAULT 0 COMMENT '第三方SKU/套餐/规格_ID',
  `name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `stock` int(11) NOT NULL DEFAULT 0 COMMENT '库存',
  `sales` int(11) NOT NULL DEFAULT 0 COMMENT '销售量',
  `stock_frozen` int(11) NOT NULL DEFAULT 0 COMMENT '冻结库存，固定分配给应用的库存',
  `price` int(11) NOT NULL DEFAULT 0 COMMENT '单价（单位分）',
  `cost_price` int(11) NOT NULL DEFAULT 0 COMMENT '成本价（单位分）',
  `origin_price` int(11) NOT NULL DEFAULT 0 COMMENT '原价（单位分）',
  `single_max` int(11) NULL DEFAULT 0 COMMENT '允许单次最大购买量,0为不限',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态：0下架、1上架、2售罄',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_spu_id`(`spu_id`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE,
  INDEX `idx_createdat`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品规则表sku-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goods_spu
-- ----------------------------
DROP TABLE IF EXISTS `goods_spu`;
CREATE TABLE `goods_spu`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `supplier_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '供应商id',
  `location_id` int(11) NOT NULL DEFAULT 0 COMMENT '站点ID',
  `cate_id` int(11) NOT NULL DEFAULT 0 COMMENT '类目ID',
  `third_spu_id` int(11) NOT NULL DEFAULT 0 COMMENT '第三方商品ID',
  `name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `share_text` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分享文字介绍',
  `price` int(10) NOT NULL DEFAULT 0 COMMENT '单价（单位分）',
  `remarks` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '说明',
  `face_img` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面图片',
  `address` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商家地址',
  `tel` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商家联系电话',
  `stock` int(11) NULL DEFAULT 0 COMMENT '库存',
  `sales` int(11) NULL DEFAULT 0 COMMENT '销售量',
  `single_min` int(11) NULL DEFAULT 1 COMMENT '允许单次最小购买量',
  `single_max` int(11) NULL DEFAULT 0 COMMENT '允许单次最大购买量,0为不限',
  `begin_time` int(11) NULL DEFAULT 1 COMMENT '开始时间',
  `end_time` int(11) NULL DEFAULT 0 COMMENT '结束时间',
  `is_order_address` tinyint(1) NULL DEFAULT 0 COMMENT '下单配送地址是否必填',
  `is_order_idcard` tinyint(1) NULL DEFAULT 0 COMMENT '下单身份证是否必填',
  `is_order_use_date` tinyint(1) NULL DEFAULT 0 COMMENT '下单使用日期是否必填',
  `is_order_delivery_time` tinyint(1) NULL DEFAULT 0 COMMENT '下单配送时间是否必填',
  `is_purchase_limit` tinyint(1) NULL DEFAULT 0 COMMENT '是否限购',
  `booking_text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '预定须知',
  `detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '图文详情',
  `attention` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '下单注意事项',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态：0下架、1上架、2售罄',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_supplier_id`(`supplier_id`) USING BTREE,
  INDEX `idx_location_id`(`location_id`) USING BTREE,
  INDEX `idx_cate_id`(`cate_id`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE,
  INDEX `idx_createdat`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品表spu-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for goods_type
-- ----------------------------
DROP TABLE IF EXISTS `goods_type`;
CREATE TABLE `goods_type`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型名称',
  `code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型code',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1有效、0无效',
  `admin_id` int(11) NOT NULL DEFAULT 0 COMMENT ' 创建人ID',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品类型表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for order_order
-- ----------------------------
DROP TABLE IF EXISTS `order_order`;
CREATE TABLE `order_order`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `sys_order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品中台订单号',
  `third_order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三订单号',
  `spu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品spuid',
  `sku_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品skuid',
  `third_spu_id` int(11) NOT NULL DEFAULT 0 COMMENT '第三方SPU_ID',
  `third_sku_id` int(11) NOT NULL DEFAULT 0 COMMENT '第三方SKU_ID',
  `location_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '站点ID',
  `total_amount` int(11) NOT NULL DEFAULT 0 COMMENT '总金额',
  `quantity` int(11) NOT NULL DEFAULT 0 COMMENT '数量',
  `customer_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '客户姓名',
  `customer_phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '客户电话',
  `address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货地址',
  `id_card` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证',
  `use_date` int(10) NOT NULL DEFAULT 0 COMMENT '使用日期',
  `delivery_time` int(10) NOT NULL DEFAULT 0 COMMENT '配送时间',
  `buy_time` int(11) NOT NULL DEFAULT 0 COMMENT '购买时间',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态：0:待发货，10:已完成,20:已核销，30:已退款',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE,
  INDEX `idx_spu_id`(`spu_id`) USING BTREE,
  INDEX `idx_sku_id`(`sku_id`) USING BTREE,
  INDEX `idx_order_no`(`order_no`) USING BTREE,
  INDEX `idx_third_order_no`(`third_order_no`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE,
  INDEX `idx_createdat`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '主订单表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for order_order_address
-- ----------------------------
DROP TABLE IF EXISTS `order_order_address`;
CREATE TABLE `order_order_address`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) NULL DEFAULT 0 COMMENT '用户ID',
  `order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `country` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '国家',
  `province` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '省',
  `city` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '市',
  `address` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '地址',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '收货人',
  `tel` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '收货人电话',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `order_no`(`order_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收货地址表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for order_order_sub
-- ----------------------------
DROP TABLE IF EXISTS `order_order_sub`;
CREATE TABLE `order_order_sub`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动id',
  `order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `sub_order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '子订单号',
  `sys_sub_order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品中台子订单号',
  `third_sub_order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三方子订单号',
  `spu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品spuid',
  `sku_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品skuid',
  `third_spu_id` int(11) NOT NULL DEFAULT 0 COMMENT '第三方SPU_ID',
  `third_sku_id` int(11) NOT NULL DEFAULT 0 COMMENT '第三方SKU_ID',
  `code` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电子码',
  `qr_code_url` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '二维码链接',
  `detail_url` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单详情',
  `booking_url` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '预约链接',
  `qr_code_img` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '二维码图片地址',
  `consume_time` int(10) NOT NULL DEFAULT 0 COMMENT '核销时间',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态：0:待发货，10:已完成,20:已核销，30:已退款',
  `is_del` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除：1是0否',
  `created_at` int(11) NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE,
  INDEX `idx_spu_id`(`spu_id`) USING BTREE,
  INDEX `idx_sku_id`(`sku_id`) USING BTREE,
  INDEX `idx_order_no`(`order_no`) USING BTREE,
  INDEX `idx_is_del`(`is_del`) USING BTREE,
  INDEX `idx_createdat`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '子订单表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_user
-- ----------------------------
DROP TABLE IF EXISTS `user_user`;
CREATE TABLE `user_user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `mobile` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '手机号',
  `unique_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '唯一标识(openid客户ID)',
  `fullname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatarurl` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别：0未知1男2女',
  `country` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '国家',
  `province` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '省',
  `city` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '市',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `mobile`(`mobile`) USING BTREE,
  INDEX `unique_id`(`unique_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表-txbao' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_user_address
-- ----------------------------
DROP TABLE IF EXISTS `user_user_address`;
CREATE TABLE `user_user_address`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) NULL DEFAULT 0 COMMENT '用户ID',
  `tag` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '标签',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '收货人',
  `tel` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '收货人电话',
  `country` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '国家',
  `province` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '省',
  `city` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '市',
  `address` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '地址',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收货地址表-txbao' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
