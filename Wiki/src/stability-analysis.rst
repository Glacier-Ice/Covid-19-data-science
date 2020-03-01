.. _stability-analysis:

Stability Analysis
================================================

Given a system of first-order differential equations

.. math::
   \frac{d \mathbf{x}(t)}{dt} = \mathbf{f}(\mathbf x(t)),

the equilibrium points are determined by :math:`\frac{d \mathbf{x}(t)}{dt} = 0`, i.e.,

.. math::
   \mathbf{f}(\mathbf x(t)) = 0,

from which we solve the equilibrium point :math:`\mathbf x(t) = \mathbf x_0`

.. _linear-stability-analysis:

Linear Stability Analysis
--------------------------------

The system of differential equations is linearized using perturbations. For simplicity, we consider the following two-dimensional system,

.. math::
   \frac{x_1(t)}{dt} &= f_1(x_1, x_2)  \\
   \frac{x_2(t)}{dt} &= f_2(x_1, x_2).
   :label: eqn-sa-2-d-1ode

The equilibrium point :math:`x_{1,0}, x_{2,0}` is determined by

.. math::
   f_1(x_{1,0}, x_{2,0}) &= 0 \\
   f_2(x_{1,0}, x_{2,0}) &= 0.

We perturb the equations around its equilibrium point,

.. math::
   x_1(t) &\to x_{1,0} + \delta x_1(t) \\
   x_2(t) &\to x_{2,0} + \delta x_2(t),

where :math:`\delta x_1(t)` and :math:`\delta x_2(t)` are small quantities.

The equation :eq:`eqn-sa-2-d-1ode` becomes

.. math::
   \frac{d\delta x_1(t)}{dt} &= f_1(x_{1,0} + \delta x_1(t), x_{2,0} + \delta x_2(t)) \\
   \frac{d\delta x_2(t)}{dt} &= f_2(x_{1,0} + \delta x_1(t), x_{2,0} + \delta x_2(t)).
   :label: eqn-sa-2-d-1ode-perturbation

We apply Talor series to :math:`f_1(x_{1,0} + \delta x_1(t), x_{2,0} + \delta x_2(t))`,

.. math::
   &f_1(x_{1,0} + \delta x_1(t), x_{2,0} + \delta x_2(t)) \\
   =&  \left.\frac{d f_1(x_1, x_2)}{d x_1} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}} \delta x_1 + \left.\frac{d f_1(x_1, x_2)}{d x_2} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}} \delta x_2 + \mathscr{O},

where we have used :math:`f_1(x_{1,0}, x_{2,0}) = 0`. The perturbed equation :eq:`eqn-sa-2-d-1ode-perturbation` is linearized

.. math::
   \frac{d\delta x_1(t)}{dt} &=  \left.\frac{d f_1(x_1, x_2)}{d x_1} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}} \delta x_1 + \left.\frac{d f_1(x_1, x_2)}{d x_2} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}} \delta x_2 \\
   \frac{d\delta x_2(t)}{dt} &=  \left.\frac{d f_2(x_1, x_2)}{d x_1} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}} \delta x_1 + \left.\frac{d f_2(x_1, x_2)}{d x_2} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}} \delta x_2.
   :label: eqn-sa-2-d-1ode-perturbation-linearized

The equation :eq:`eqn-sa-2-d-1ode-perturbation-linearized` can be rewritten into a matrix form

.. math::
   \frac{d}{dt} \begin{pmatrix}
   \delta x_1(t) \\
   \delta x_2(t)
   \end{pmatrix} = \left. \begin{pmatrix}
   \frac{d f_1(x_1, x_2)}{d x_1}  & \frac{d f_1(x_1, x_2)}{d x_2} \\
   \frac{d f_2(x_1, x_2)}{d x_1} & \frac{d f_2(x_1, x_2)}{d x_2}
   \end{pmatrix} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}}\begin{pmatrix}
   \delta x_1(t) \\
   \delta x_2(t)
   \end{pmatrix}.
   :label: eqn-sa-2-d-1ode-perturbation-linearized-mat

For simplicity, we define the Jacobian matrix

.. math::
   \mathbf J = \left. \begin{pmatrix}
   \frac{d f_1(x_1, x_2)}{d x_1}  & \frac{d f_1(x_1, x_2)}{d x_2} \\
   \frac{d f_2(x_1, x_2)}{d x_1} & \frac{d f_2(x_1, x_2)}{d x_2}
   \end{pmatrix} \right\vert_{x_1=x_{1,0}, x_2=x_{2,0}},

so that the matrix form of the linearized equation :eq:`eqn-sa-2-d-1ode-perturbation-linearized-mat` becomes,

.. math::
   \frac{d}{dt} \begin{pmatrix}
   \delta x_1(t) \\
   \delta x_2(t)
   \end{pmatrix} = \mathbf J \begin{pmatrix}
   \delta x_1(t) \\
   \delta x_2(t)
   \end{pmatrix}.

To investigate the stability of the differential system, we assume that

.. math::
   \begin{pmatrix}
   \delta x_1(t) \\
   \delta x_2(t)
   \end{pmatrix} = \begin{pmatrix}
   \delta x_1(t_0) \\
   \delta x_2(t_0)
   \end{pmatrix} e^{\lambda t},

which leads to linear equations

.. math::
   \begin{pmatrix}
   \delta x_1(t_0) \\
   \delta x_2(t_0)
   \end{pmatrix} \lambda = \mathbf J \begin{pmatrix}
   \delta x_1(t_0) \\
   \delta x_2(t_0)
   \end{pmatrix}.

For non-trivial solutions, we require

.. math::
   \operatorname{Det}(\mathbf J - \lambda \mathbf I) = 0,

which is also the eigen value problem of the Jacobian matrix. We expand the determinant

.. math::
   \lambda^2 - (J_{11} + J_{22}) \lambda + (J_{11}J_{22} - J_{12}J_{21}) = 0.
   :label: eqn-sa-2-d-1ode-perturbation-linearized-det

For real positive solutions :math:`\lambda>0`, we get an exponential growing result for the linearized equation. Any deviation from the equilibrium point leads to a run-away process and the system moves further away from the equilibrium point. For real negative solutions :math:`\lambda < 0`, the system will move back to the equilibrium point given any deviations from the equilibrium. Imaginary solutions of :math:`\lambda` leads to oscillations.

