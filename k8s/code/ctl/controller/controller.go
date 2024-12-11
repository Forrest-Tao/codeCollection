package controllers

import (
	"context"
	"fmt"
	"forrest/codeCollection/k8s/code/ctl/apis"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type SimpleReconciler struct {
	client.Client
}

// Reconcile 是控制器的核心逻辑
func (r *SimpleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// 获取资源
	simple := &apis.Simple{}
	if err := r.Get(ctx, req.NamespacedName, simple); err != nil {
		log.Error(err, "unable to fetch Simple")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// 打印资源的 message
	log.Info("Reconciling Simple", "message", simple.Spec.Message)

	// Example: 为资源增加一个注解
	if simple.Annotations == nil {
		simple.Annotations = make(map[string]string)
	}
	simple.Annotations["reconciled-at"] = fmt.Sprintf("%v", ctrl.Now())

	// 更新资源
	if err := r.Update(ctx, simple); err != nil {
		log.Error(err, "unable to update Simple")
		return ctrl.Result{}, err
	}

	// Example: 创建或清理相关资源
	controllerutil.AddFinalizer(simple, "example.com/finalizer")

	// 返回结果
	return ctrl.Result{}, nil
}

// SetupWithManager 绑定控制器到 Manager
func (r *SimpleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apis.Simple{}). // 监听 Simple 资源
		Complete(r)
}
